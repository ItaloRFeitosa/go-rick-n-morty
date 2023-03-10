# š Go Rick Nā Morty

## š Description

Web API to apply some design patterns

## š» Functionalities

- [x]  Search characters from the Rick n Morty animation

## š Patterns applied so far

- [x]  Creational
    - [x]  Singleton
- [x]  Structural
    - [x]  Decorator
    - [x]  Proxy
    - [x]  Adapter

## šØ Installation

1. Create .env
2. `go mod tidy`
3. `go run ./cmd/api/main.go`

Note: for now we do not have a docker-compose to make it easier to use `redis`, but here is a [tutorial](https://medium.com/@prog.tiago/redis-instalando-via-docker-58cb1d2cfb3b) to bring up a redis instance with docker or, if you prefer, you can create an instance on [upstash](https://webhook.site/86bea15e-e89e-4d55-8ec7-e425630a50b0), it is very easy and fast to bring up.

## š ļø Tech Stack

- go 1.18
- viper
- go-query-string
- fiber
- redis
- go-cache
- go-redis
