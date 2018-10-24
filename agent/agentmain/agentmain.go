/*
Copyright 2017 Google Inc. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"hash/fnv"
	"net/http"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/statslog"
	"github.com/GoogleCloudPlatform/cloud-ingest/gcloud"
	"github.com/golang/glog"
	"google.golang.org/api/option"
)

const (
	listProgressTopic = "cloud-ingest-list-progress"
	copyProgressTopic = "cloud-ingest-copy-progress"
	pulseTopic        = "cloud-ingest-pulse"
	controlTopic      = "cloud-ingest-control"

	listSubscription    = "cloud-ingest-list"
	copySubscription    = "cloud-ingest-copy"
	controlSubscription = "cloud-ingest-control"
)

var (
	projectID                string
	numberThreads            int
	maxPubSubLeaseExtenstion time.Duration
	credsFile                string
	chunkSize                int

	pubsubPrefix string

	skipProcessListTasks bool
	skipProcessCopyTasks bool
	printVersion         bool

	pulseFrequency int
	pulseRun       bool //used to start or stop pulse

	enableStatsLog bool

	// Fields used to display version information. These defaults are
	// overridden when the release script builds in values through ldflags.
	buildVersion = "v0.0.0-development"
	buildCommit  = "(local)"
	buildDate    = "Unknown; this is not an official release."
)

func init() {
	flag.StringVar(&projectID, "projectid", "",
		"The Pub/Sub topics and subscriptions project id. Must be set!")
	flag.StringVar(&credsFile, "creds-file", "",
		"The service account JSON key file. Use the default credentials if empty.")
	flag.IntVar(&chunkSize, "chunk-size", 1<<25,
		"The resumable upload chuck size, default 32MB.")
	flag.IntVar(&numberThreads, "threads", 100,
		"The number of threads to process the copy tasks. If 0, will use the "+
			"default Pub/Sub client value (1000).")
	flag.DurationVar(&maxPubSubLeaseExtenstion, "pubsub-lease-extension", 0,
		"The max duration to extend the leases for a Pub/Sub message. If 0, will "+
			"use the default Pub/Sub client value (10 mins).")
	flag.Uint64Var(&glog.MaxSize, "max-log-size", 1<<28,
		"The maximum size of a log file in bytes, default 268435456 bytes(256MB).")

	flag.BoolVar(&skipProcessListTasks, "skip-list", false,
		"Skip processing list tasks.")
	flag.BoolVar(&skipProcessCopyTasks, "skip-copy", false,
		"Skip processing copy tasks.")
	flag.BoolVar(&printVersion, "version", false,
		"Print build/version info and exit.")

	flag.StringVar(&pubsubPrefix, "pubsub-prefix", "",
		"Prefix of Pub/Sub topics and subscriptions names.")

	flag.IntVar(&pulseFrequency, "pulse-frequency", 10, "the number of seconds the agent will wait before sending a pulse")
	flag.BoolVar(&pulseRun, "pulse-run", true, "Send pulse")

	flag.BoolVar(&agent.ControlEnabled, "enable-agent-control", true, "Enable the agent to be controlled by the service backend.")
	flag.BoolVar(&enableStatsLog, "enable-stats-log", true, "Enable stats logging to INFO logs.")

	flag.Parse()
}

func printVersionInfo() {
	fmt.Printf("Google Cloud Ingest Agent %s\n", buildVersion)
	fmt.Printf("Git Commit: %s\nBuild Date: %s\n", buildCommit, buildDate)
}

// waitOnSubscription blocks until either the passed-in subscription exists, or
// an error occurs (including context end). In all cases where we return without
// the subscription existing, we return the relevant error.
func waitOnSubscription(ctx context.Context, sub *pubsub.Subscription) error {
	exists, err := sub.Exists(ctx)
	if err != nil {
		return err
	}

	t := time.NewTicker(10 * time.Second)

	for !exists {
		fmt.Printf("Waiting for subscription %s to exist. If this is the first run for this project, "+
			"create your first transfer job.\n", sub.String())

		select {
		case <-t.C:
			exists, err = sub.Exists(ctx)
			if err != nil {
				t.Stop()
				return err
			}
		case <-ctx.Done():
			t.Stop()
			return ctx.Err()
		}
	}

	t.Stop()
	fmt.Printf("Subscription %s is ready!\n", sub.String())
	return nil
}

// waitOnTopic blocks until either the passed-in PSTopic exists, or
// an error occurs (including context end). In all cases where we return without
// the PSTopic existing, we return the relevant error.
func waitOnTopic(ctx context.Context, topic gcloud.PSTopic) error {
	exists, err := topic.Exists(ctx)
	if err != nil {
		return err
	}

	t := time.NewTicker(10 * time.Second)
	for !exists {
		fmt.Printf("Waiting for topic %s to exist.\n", topic.ID())

		select {
		case <-t.C:
			exists, err = topic.Exists(ctx)
			if err != nil {
				t.Stop()
				return err
			}
		case <-ctx.Done():
			t.Stop()
			return ctx.Err()
		}
	}

	t.Stop()
	fmt.Printf("Topic %s is ready!\n", topic.ID())
	return nil
}

func subscribeToControlTopic(ctx context.Context, client *pubsub.Client, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	hostname, err := agent.GetHostName()
	if err != nil {
		return nil, err
	}
	h := fnv.New64a()
	h.Write([]byte(hostname))
	h.Write([]byte(agent.GetProcessId()))

	subID := fmt.Sprintf("%s%s-%d", pubsubPrefix, controlSubscription, h.Sum64())
	sub := client.Subscription(subID)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		glog.Infof("Subscription %s already exists, probably another agent created it before.", sub.String())
		return sub, nil
	}
	return client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{Topic: topic})
}

func main() {
	defer glog.Flush()
	ctx := context.Background()

	if printVersion {
		printVersionInfo()
		os.Exit(0)
	}

	var pubSubErr, storageErr, httpcErr error
	var pubSubClient *pubsub.Client
	var storageClient *storage.Client
	var httpc *http.Client

	if credsFile != "" {
		clientOptions := option.WithCredentialsFile(credsFile)
		pubSubClient, pubSubErr = pubsub.NewClient(ctx, projectID, clientOptions)
		storageClient, storageErr = storage.NewClient(ctx, clientOptions)
		httpc, httpcErr = agent.NewResumableHttpClient(ctx, clientOptions)
	} else {
		pubSubClient, pubSubErr = pubsub.NewClient(ctx, projectID)
		storageClient, storageErr = storage.NewClient(ctx)
		httpc, httpcErr = agent.NewResumableHttpClient(ctx)
	}

	if pubSubErr != nil {
		glog.Fatalf("Can't create Pub/Sub client, error: %+v\n", pubSubErr)
	}
	if storageErr != nil {
		glog.Fatalf("Can't create storage client, error: %+v\n", storageErr)
	}
	if httpcErr != nil {
		glog.Fatalf("Can't create http.Client, error: %+v\n", httpcErr)
	}

	var wg sync.WaitGroup

	if pulseRun {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pulseTopicWrapper := gcloud.NewPubSubTopicWrapper(pubSubClient.Topic(pubsubPrefix + pulseTopic))

			// Wait for pulse topic to exist.
			if err := waitOnTopic(ctx, pulseTopicWrapper); err != nil {
				glog.Fatalf("Could not get PulseTopic: %s \n error: %v ", pulseTopicWrapper.ID(), err)
			}

			ph, err := agent.NewPulseHandler(pulseTopicWrapper, int32(pulseFrequency))
			if err != nil {
				glog.Fatalf("Could not create a PulseHandler with Topic: %v and Frequency: %v \n error: %v ", pulseTopicWrapper.ID(), pulseFrequency, err)
			}

			ph.Run(ctx)
		}()
	}

	var sl *statslog.StatsLog
	if enableStatsLog {
		sl = statslog.New()
		go sl.PeriodicallyLogStats(ctx)
	}

	if !skipProcessListTasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			listSub := pubSubClient.Subscription(pubsubPrefix + listSubscription)
			listSub.ReceiveSettings.MaxExtension = maxPubSubLeaseExtenstion
			listSub.ReceiveSettings.MaxOutstandingMessages = numberThreads
			listTopic := pubSubClient.Topic(pubsubPrefix + listProgressTopic)
			listTopicWrapper := gcloud.NewPubSubTopicWrapper(listTopic)

			// Wait for list subscription to exist.
			if err := waitOnSubscription(ctx, listSub); err != nil {
				glog.Fatalf("Could not find list subscription %s, error %+v", listSub.String(), err)
			}

			// Wait for list topic to exist.
			if err := waitOnTopic(ctx, listTopicWrapper); err != nil {
				glog.Fatalf("Could not find list topic %s, error %+v", listTopicWrapper.ID(), err)
			}

			listProcessor := agent.WorkProcessor{
				WorkSub:       listSub,
				ProgressTopic: listTopic,
				Handler:       agent.NewListHandler(storageClient, chunkSize),
				StatsLog:      sl,
			}
			listProcessor.Process(ctx)
		}()
	}

	if !skipProcessCopyTasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			copySub := pubSubClient.Subscription(pubsubPrefix + copySubscription)
			copySub.ReceiveSettings.MaxExtension = maxPubSubLeaseExtenstion
			copySub.ReceiveSettings.MaxOutstandingMessages = numberThreads
			copyTopic := pubSubClient.Topic(pubsubPrefix + copyProgressTopic)
			copyTopicWrapper := gcloud.NewPubSubTopicWrapper(copyTopic)

			// Wait for copy subscription to exist.
			if err := waitOnSubscription(ctx, copySub); err != nil {
				glog.Fatalf("Could not find copy subscription %s, error %+v", copySub.String(), err)
			}

			// Wait for copy topic to exist.
			if err := waitOnTopic(ctx, copyTopicWrapper); err != nil {
				glog.Fatalf("Could not find copy topic %s, error %+v", copyTopicWrapper.ID(), err)
			}

			copyProcessor := agent.WorkProcessor{
				WorkSub:       copySub,
				ProgressTopic: copyTopic,
				Handler:       agent.NewCopyHandler(storageClient, chunkSize, httpc),
				StatsLog:      sl,
			}

			copyProcessor.Process(ctx)
		}()
	}

	if agent.ControlEnabled {
		if !pulseRun {
			glog.Fatalf("Command line flag 'pulse-run' should be enabled.")
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			controlTopic := pubSubClient.Topic(pubsubPrefix + controlTopic)
			controlTopicWrapper := gcloud.NewPubSubTopicWrapper(controlTopic)

			// Wait for copy topic to exist.
			if err := waitOnTopic(ctx, controlTopicWrapper); err != nil {
				glog.Fatalf("Could not find control topic %s, error %+v", controlTopicWrapper.ID(), err)
			}

			controlSub, err := subscribeToControlTopic(ctx, pubSubClient, controlTopic)
			if err != nil {
				glog.Fatalf("Could not create subscription to control topic %v, with err: %v", controlTopic, err)
			}

			ch := agent.NewControlHandler(controlSub)
			if err := ch.HandleControlMessages(ctx); err != nil {
				glog.Fatalf("Failed handling control messages with err: %v.", err)
			}
		}()
	}

	wg.Wait()
}
