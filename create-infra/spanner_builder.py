# -*- coding: utf-8 -*-
# Copyright 2017 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Google Cloud Spanner admin utilities."""

from google.cloud import spanner


class SpannerBuilder(object):
  """Manipulates creation/deletion of spanner instances/databases."""

  def __init__(self,
               instance_id,
               location='us-central1',
               node_count=1):
    client = spanner.Client()
    # Search for a configuration matching the input location.
    config = None
    configs = client.list_instance_configs()
    for c in configs:
      if c.name.endswith(location):
        config = c
        break
    if not config:
      raise Exception(
          'Can not get spanner config for location {}.'.format(location))
    self.instance = client.instance(
        instance_id, configuration_name=config.name, node_count=node_count,
        display_name=instance_id)

  def CreateInstance(self):
    """Creates the spanner instance."""
    operation = self.instance.create()
    operation.result()

  def DeleteInstance(self):
    """Deletes the spanner instance."""
    self.instance.delete()

  def CreateDatabase(self, database_id, schema_file):
    """Creates a database in the spanner instance."""
    # Read the schema file
    with open(schema_file, 'r') as f:
      schema = f.read()

    # Split the schema into statement array.
    statements = [x for x in schema.split('\n\n') if x]

    database = self.instance.database(database_id, ddl_statements=statements)
    operation = database.create()
    operation.result()
