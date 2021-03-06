# Kubemq GCP Cloud Functions Target Connector

Kubemq gcp Cloud Functions target connector allows services using kubemq server to access call a function target on gcp Cloud Function.
The connector synchronously invokes a deployed Cloud Function, very limited traffic is allowed 16 per 100 seconds.

## Prerequisites
The following are required to run the redis target connector:

- kubemq cluster
- Gcp Cloud function targe.
- kubemq-targets deployment

## Configuration

Kafka source connector configuration properties:

| Properties Key | Required | Description                                    | Example          |
|:---------------|:---------|:-----------------------------------------------|:-----------------|
| project        | yes      | gcp project name                               | "testproject"    |
| credentials    | yes      | gcp service account key location (json)        | "TestTopic"      |
| location_match  | no       | match missing function location (default true) | "false"          |


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
      kind: target.gcp.cloudfunctions
      name: gcp-functions-target
      properties:
        project: "testproject"
        credentials: "./../key.json"
```

## Usage

### Get Request

Get request metadata setting:

| Metadata Key | Required | Description                                          | Possible values                         |
|:-------------|:---------|:-----------------------------------------------------|:----------------------------------------|
| name         | yes      | name of the gcp function                             | "test_function"                          |
| project      | no       | gcp project name (override parent value)             | "testproject1"                          |
| location     | no       | gcp project location (missing value will be matched) | "us-central1"                           |


Example:

```json
{
  "metadata": {
    "name": "test_function",
    "location": "us-central1"    
  },
  "data": "eyJtZXNzYWdlIjoidGVzdCJ9"
}
```
