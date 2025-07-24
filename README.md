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

## Database management

This application uses sqlite3, install if you dont have.

```shell
sudo apt install sqlite3
```

Then execute sqlite prompt in db file

```shell
sqlite3 data.db
```

## Query samples for GraphQL playgroung

Select categories

```json
query queryCategories {
  categories {
    id
    name
  }
}
```

Create new category

```json
mutation createCategory {
  createCategory(input: {name: "Blue Cheese 250gr"}) {
    id
  }
}
```
