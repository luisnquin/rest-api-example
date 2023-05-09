
# Blind creator REST API test

## Requirements

- REST API(don't use frameworks)
- Use of GORM with PostgreSQL or MySQL
- Deployed in AWS Lambda or AWS EC2
- Must have a pagination endpoint with relationships.

## Tools

### Required

- `go >=1.20`
- [just](https://github.com/casey/just)
- [docker](https://docs.docker.com/get-docker)
- [docker-compose](https://docs.docker.com/compose/install)
- [air](https://github.com/cosmtrek/air) (For development)

### Recommended

- [golangci-lint](https://golangci-lint.run/usage/install)
- [panicparse](https://github.com/maruel/panicparse)

## How to run

Make sure you have correctly installed the tools from the previous section

```bash
# If first time
$ just migrate

# Start the server
$ just start
```
