
# Blind creator REST API test

## Requirements

- REST API(don't use frameworks)
- Use of GORM with PostgreSQL or MySQL
- Deployed in AWS Lambda or AWS EC2
- Must have a pagination endpoint with relationships.

## Tools

### Required

- `go >=1.20`
- `just`
- `docker`
- `docker-compose`

### Recommended

- `golangci-lint`

## How to run

Make sure you have correctly installed the tools from the previous section

```bash
# If first time
$ just migrate

# Start the server
$ just start
```
