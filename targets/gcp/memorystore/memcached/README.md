# Kubemq Memcached-gcp Target Connector

Kubemq memcached target connector allows services using kubemq server to access memcached server functions such `set`, `get` and `delete`.

## Prerequisites
The following are required to run the memcached target connector:

- kubemq cluster
- memcached server
- access to gcp Memcached server
- kubemq-targets deployment

## Configuration

Memcached target connector configuration properties:

| Properties Key          | Required | Description                                       | Example                           |
|:------------------------|:---------|:--------------------------------------------------|:----------------------------------|
| hosts                   | yes      | memcached servers list address separated by comma | "localhost:11211,localhost:11212" |
| max_idle_connections    | no       | set max idle connection                           | "2"                               |
| default_timeout_seconds | no       | set default timeout seconds for operation         | "10"                              |

Example:

```yaml
bindings:
  - name: kubemq-query-gcp-memcached
    source:
      kind: source.query
      name: kubemq-query
      properties:
        host: "localhost"
        port: "50000"
        client_id: "kubemq-query-gcp-memcached-connector"
        auth_token: ""
        channel: "query.memcached"
        group:   ""
        concurrency: "1"
        auto_reconnect: "true"
        reconnect_interval_seconds: "1"
        max_reconnects: "0"
    target:
      kind: target.gcp.cache.memcached
      name: target-gcp-memcached
      properties:
        hosts: "localhost:11211"
        max_idle_connections: "2"
        default_timeout_seconds: "10"
```

## Usage

### Get Request

Get request metadata setting:

| Metadata Key | Required | Description      | Possible values |
|:-------------|:---------|:-----------------|:----------------|
| key          | yes      | memcached key string | any string      |
| method       | yes      | get              | "get"           |


Example:

```json
{
  "metadata": {
    "key": "your-memcached-key",
    "method": "get"
  },
  "data": null
}
```

### Set Request

Set request metadata setting:

| Metadata Key | Required | Description      | Possible values |
|:-------------|:---------|:-----------------|:----------------|
| key          | yes      | memcached key string | any string      |
| method       | yes      | set              | "set"           |

Set request data setting:

| Data Key | Required | Description                   | Possible values     |
|:---------|:---------|:------------------------------|:--------------------|
| data     | yes      | data to set for the memcached key | base64 bytes array |

Example:

```json
{
  "metadata": {
    "key": "your-memcached-key",
    "method": "set"
  },
  "data": "c29tZS1kYXRh" 
}
```
### Delete Request

Delete request metadata setting:

| Metadata Key | Required | Description      | Possible values |
|:-------------|:---------|:-----------------|:----------------|
| key          | yes      | memcached key string | any string      |
| method       | yes      | delete           | "delete"        |


Example:

```json
{
  "metadata": {
    "key": "your-memcached-key",
    "method": "delete"
  },
  "data": null
}
```
