# sternx_challenge

This task involves writing a SQL query to retrieve the latest trade for each symbol from a database schema and creating a microservice with an HTTP API that executes the query and returns the results in JSON format, with an emphasis on code readability and performance optimization.

## Architecture

![Alt text](api-docs/images/architecture.png "Architect")


## Development

To start the sternx-challenge in the run:

```
go run cmd/main.go
```
To start the sternxcli-challenge in the run:

```
go run cmd/cli/main.go -num-record <record-numbers>
```
### Getting the last trades for each Instrument 
```
curl -XGET https://localhost:8088/api/v1/latest-trades
```
