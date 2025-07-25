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
## Database structure

In case of lost your data.db, re-create then with this structure:

```sql
create table categories (id string, name string);
create table products (id string, name string, description string, category_id string);
```

## Query samples for GraphQL playgroung

Select categories

```graphql
query queryCategories {
  categories {
    id
    name
  }
}
```
Select categories and products

```graphql
query queryCategoriesProducts {
  categories {
    id
    name
    products {
      id
      name
    }
  }
}
```

Create new category

```graphql
mutation createCategory {
  createCategory(input: {name: "Blue Cheese"}) {
    id
  }
}
```

Select products

```graphql
query queryProducts {
  products {
    id
    name
    description
  }
}
```

Select products with category

```graphql
query queryProductsCategory {
  products {
    id
    name
    description
    category {
      id
      name
    }
  }
}
```

Create new product

```graphql
mutation createProduct {
  createProduct(input: {
    name: "Vigor Gorgon Space",
    description: "Vigor Gorgonzola 250gr in a tablet",
    categoryId: "ce9b0a14-32f6-47fc-93da-245a6da05b6d"
  }) {
    id
  }
}
```
