CREATE TABLE JobConfigs (
     JobConfigId STRING(MAX) NOT NULL,
     JobSpec STRING(MAX),
   ) PRIMARY KEY(JobConfigId)

CREATE TABLE JobRuns (
     JobConfigId STRING(MAX) NOT NULL,
     JobRunId STRING(MAX) NOT NULL,
     Status INT64,
     JobCreationTime INT64,
   ) PRIMARY KEY(JobConfigId, JobRunId),
   INTERLEAVE IN PARENT JobConfigs ON DELETE NO ACTION

CREATE TABLE Tasks (
      JobConfigId STRING(MAX) NOT NULL,
      JobRunId STRING(MAX) NOT NULL,
      TaskId STRING(MAX) NOT NULL,
      TaskSpec STRING(MAX) NOT NULL,
      Status INT64 NOT NULL,
      WorkerId STRING(MAX),
      LastModificationTime INT64,
      FailureMessage STRING(MAX),
    ) PRIMARY KEY(JobConfigId, JobRunId, TaskId),
    INTERLEAVE IN PARENT JobRuns ON DELETE NO ACTION

 CREATE INDEX TasksByStatus ON Tasks(Status)