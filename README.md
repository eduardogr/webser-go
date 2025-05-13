[![Go Report Card](https://goreportcard.com/badge/github.com/eduardogr/webser-go)](https://goreportcard.com/report/github.com/eduardogr/webser-go)

<h1 align="center"> go-webserver </h1> <br>

## Bootstrapping go-server

### Installations needed

- go
- docker
- docker-compose

### Running

- make wire
- make build
- make up
  - verify you are running this with proper permissions

## exposed API

```
/api/v0
Do nothing now, just to be an example for a multiple api/vX format in router.go file

GET /api/v1/numbers, return all numbers stored

POST /api/v1/numbers -d '{"ID":11}', creating a number in the storage
```
