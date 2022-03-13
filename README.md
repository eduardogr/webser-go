<h1 align="center"> webserver-go </h1> <br>

## Bootstrapping web-server-go

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

## TODO

- [ ] Migrations of the DDBB schema
- [ ] Implement the concept of external provider for secrets
- [ ] Complete API documentation for clients' usage
- [ ] Improve docker image building for go mod download slow down
- [ ] Testing when initializing database schema did not work
- [ ] Adding tests for api specifications. Package pkg.server
- [ ] Adding tests for pkg.repository package
- [ ] Adding tests for pkg.api package
