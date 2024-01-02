
# Server example

## Tools

### Required

- `go >=1.20`
- `git`
- [just](https://github.com/casey/just)
- [docker](https://docs.docker.com/get-docker)
- [docker-compose](https://docs.docker.com/compose/install)
- [air](https://github.com/cosmtrek/air) (only for development)

### Recommended

- [golangci-lint](https://golangci-lint.run/usage/install)
- [panicparse](https://github.com/maruel/panicparse)
- [pg-ping](https://github.com/luisnquin/pg-ping)

## Get started

- Ensure that you have all the `required tools` of the previous section
- Clone the project repository and stay in the root directory

    ```bash
    git clone https://github.com/luisnquin/server-example.git /tmp/luisnquin-server-example
    cd /tmp/luisnquin-server-example
    ```

- Copy the `.env.example` file and rename it to `.env`, otherwise you'll not be
able to do the next steps.

- You will need to create the database and migrate all the database tables with
mock data, for that just execute:

    ```bash
    # Erases, starts and migrates the database
    $ just migrate
    ```

### After that

- If is your first computer startup or if the services were killed then execute:

    ```bash
    # The docker compose services will need to be started
    $ just compose-up
    ```

- Just start the server:

    ```bash
    # Compiles and executes the server
    $ just start
    ```

- Start the server with live reload:

    ```bash
    # Uses `air` by behind to provide live-reload
    $ just dev
    ```

## Requirements

- REST API(don't use frameworks)
- Use of GORM with PostgreSQL or MySQL
- Deployed in AWS Lambda or AWS EC2
- Must have a pagination endpoint with relationships.
