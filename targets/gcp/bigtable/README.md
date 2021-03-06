# Kubemq bigtable target Connector

Kubemq gcp-bigtable target connector allows services using kubemq server to access google bigtable server.

## Prerequisites
The following required to run the gcp-bigtable target connector:

- kubemq cluster
- gcp-bigtable set up
- kubemq-source deployment

## Configuration

bigtable target connector configuration properties:

| Properties Key | Required | Description                                | Example                     |
|:---------------|:---------|:-------------------------------------------|:----------------------------|
| project_id     | yes      | gcp bigtable project_id                    | "<googleurl>/myproject"     |
| credentials    | yes      | gcp credentials files                      | "<google json credentials"  |
| instance       | yes      | bigtable instance name                     | "<bigtable instance name"   |


Example:

```yaml
bindings:
  - name: kubemq-query-gcp-bigtable
    source:
      kind: source.query
      name: kubemq-query
      properties:
        host: "localhost"
        port: "50000"
        client_id: "kubemq-query-gcp-bigtable-connector"
        auth_token: ""
        channel: "query.gcp.bigtable"
        group:   ""
        concurrency: "1"
        auto_reconnect: "true"
        reconnect_interval_seconds: "1"
        max_reconnects: "0"
    target:
      kind: target.gcp.bigtable
      name: target-gcp-bigtable
      properties:
        project_id: "id"
        credentials: 'json'
        instance:  "instance"
```

## Usage

### Create Column Family

Create Column Family:

| Metadata Key      | Required | Description                             | Possible values                            |
|:------------------|:---------|:----------------------------------------|:-------------------------------------------|
| method            | yes      | type of method                          | "create_column_family"                     |
| column_family     | yes      | the column_family to create             | "valid unique string"                      |
| table_name        | yes      | the table name                          | "table name to assign the column family"   |


Example:

```json
{
  "metadata": {
    "method": "create_column_family",
    "column_family": "valid_unique_string",
    "table_name": "valid_table_string"
  },
  "data": null
}
```

### Create Or Delete Table

Create or delete from the table:

| Metadata Key | Required | Description                             | Possible values                         |
|:-------------|:---------|:----------------------------------------|:----------------------------------------|
| method       | yes      | type of method                          | "create_table", "delete_table"          |
| table_name   | yes      | the table name                          | "table name to delete or create"        |


Example:

```json
{
  "metadata": {
    "method": "delete_table",
    "table_name": "valid_table_string"
  },
  "data": null
}
```


### Write Rows

Write rows to table:

| Metadata Key      | Required | Description                             | Possible values                         |
|:------------------|:---------|:----------------------------------------|:----------------------------------------|
| method            | yes      | type of method                          | "write"                                 |
| table_name        | yes      | the table name                          | "table name to delete or create"        |
| column_family     | yes      | the column_family to create             | "valid unique string"                   |

Example:

```json
{
  "metadata": {
    "method": "write",
    "column_family": "valid_unique_string",
    "table_name": "valid_table_string"
  },
  "data": "eyJpZCI6MSwibmFtZSI6InRlc3QxIiwic2V0X3Jvd19rZXkiOiIxIn0="
}
```

### Delete Rows

Delete rows from the table:

| Metadata Key      | Required | Description                             | Possible values                         |
|:------------------|:---------|:----------------------------------------|:----------------------------------------|
| method            | yes      | type of method                          | "delete_row"                            |
| table_name        | yes      | the table name                          | "table name to delete or create"        |
| row_key_prefix    | yes      | the row key                             | "valid unique string"                   |

Example:

```json
{
  "metadata": {
    "method": "delete_row",
    "row_key_prefix": "valid unique string",
    "table_name": "valid_table_string"
  },
  "data": null
}
```


### Delete Rows

Delete rows from the table:

| Metadata Key      | Required | Description                  | Possible values                            |
|:------------------|:---------|:-----------------------------|:-------------------------------------------|
| method            | yes      | type of method               | "delete_row"                               |
| table_name        | yes      | the table name               | "table name to delete or create"           |
| row_key_prefix    | no       | the row key                  | "valid unique string"                      |
| column_name       | no       | the column to return         | "column name"                              |

Example:

```json
{
  "metadata": {
    "method": "get_all_rows",
    "table_name": "valid_table_string"
  },
  "data": null
}