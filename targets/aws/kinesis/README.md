# Kubemq aws kinesis Data Streams Connector

Kubemq aws-kinesis Data Streams target connector allows services using kubemq server to create and put messages on kinesis Data Streams.

## Prerequisites
The following required to run the gcp-firebase target connector:

- kubemq cluster
- aws user with kinesis Data Streams access rights
- kubemq-source deployment

## Configuration

kinesis Data Streams connector configuration properties:

| Properties Key | Required | Description                       | Example          |
|:---------------|:---------|:----------------------------------|:-----------------|
| aws_key        | yes      | aws iam  user access-keys         | * link below     |
| aws_secret_key | yes      | aws user aws_secret_key           | * link           |
| region         | no       | aws kinesis Data Streams region   | "us-east-2"      |
| token          | no       | aws user token                    |                  |

* iam examples can be found on http://docs.aws.amazon.com/console/iam/about-access-keys


Example:

```yaml
bindings:
  - name: kubemq-query-kafka
    source:
      kind: source.kubemq.query
      name: kubemq-query
      properties:
        host: "localhost"
        port: "50000"
        client_id: "kubemq-query-gcpfucntions-connector"
        auth_token: ""
        channel: "query.gcp.functions"
        group:   ""
        concurrency: "1"
        auto_reconnect: "true"
        reconnect_interval_seconds: "1"
        max_reconnects: "0"
    target:
      kind: target.aws.kenisis
      name: target.aws.kenisis
      properties:
       	aws_key: "j"
        aws_secret_key: "osuLJf+sdf"
        region: "us-east-2"
```

## Usage

### Put 

Writes a single data record into an Amazon Kinesis data stream.

Get request metadata setting:

| Metadata Key  | Required | Description                                                 | Possible values                         |
|:--------------|:---------|:------------------------------------------------------------|:----------------------------------------|
| stream        | yes      | stream name scoped by region                                | "test"                                  |
| method        | no       | request action "put" (as default) put data record to stream |"put/create"                             |
|               |          | "create" -  create a data stream                            |                                         |   
| shard_count   | no       | number of shards that the stream will use on create method  | "2"                                     |
| partition_key | no       | stream shard when put data record                           | "2"                                     | 



Example put:

```json
{
  "metadata": {
    "stream": "test",
    "method": "put",
    "shard_count":"1"
  },
  "data": "eyJtZXNzYWdlIjoidGVzdCJ9"
}
```

Example create:

```json
{
  "metadata": {
    "stream": "test",
    "method": "create",
    "partition_key":"1"
  },
  "data": null
}
```
