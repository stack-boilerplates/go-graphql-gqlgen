# go-graphql-gqlgen
Golang BFF with GraphQL provided by [gqlgen](https://gqlgen.com/) lib

## Run
```shell
go run server.go
```
Access http://localhost:8080/ for GraphQL playground

## Update GraphQL schema

When making changes in graph/schema.graphqls file, run this command to generate needed Golang models and functions

```shell
go run github.com/99designs/gqlgen generate
```
