# Unique ID Generator webservice

A Unique ID Generator exposing a REST API.

## API

GET `/v1/ids` - returns a new unique ID.

## Build and run the project

```
$ go mod tidy
$ PORT=8080 go run ./cmd/webservice/id_generator_webservice.go
```
