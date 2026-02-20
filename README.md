# graphql-books

A GraphQL API server for managing books and authors, built with Go and [gqlgen](https://github.com/99designs/gqlgen). Supports queries, mutations, and real-time subscriptions via WebSocket.

## Prerequisites

- [Go 1.25.0+](https://go.dev/dl/)

## Running

Install dependencies and start the server:

```bash
go run server.go
```

The server starts on port **8080** with two endpoints:

| Endpoint | Description |
|---|---|
| `http://localhost:8080/` | GraphQL Playground (browser IDE) |
| `http://localhost:8080/graphql` | GraphQL API |

## Testing

There are no automated tests yet. You can test the API manually using the GraphQL Playground at `http://localhost:8080/`.

### Example queries

**Get all books:**

```graphql
query {
  books {
    id
    title
    year
    author {
      name
    }
  }
}
```

**Get a single book by ID:**

```graphql
query {
  book(id: "book-1") {
    title
    year
    author {
      name
    }
  }
}
```

**Get all authors with their books:**

```graphql
query {
  authors {
    id
    name
    books {
      title
    }
  }
}
```

**Create a new book:**

```graphql
mutation {
  createBook(input: { title: "New Book", year: 2024, authorId: "author-1" }) {
    id
    title
    author {
      name
    }
  }
}
```

**Subscribe to new books (WebSocket):**

```graphql
subscription {
  bookAdded {
    id
    title
    author {
      name
    }
  }
}
```

### Testing with curl

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ books { id title year author { name } } }"}'
```

## Code Generation

If you modify the GraphQL schema (`graph/schema.graphqls`), regenerate the Go code:

```bash
go run github.com/99designs/gqlgen generate
```

## Project Structure

```
.
├── server.go                 # Entry point - HTTP server setup
├── gqlgen.yml                # gqlgen code generation config
├── go.mod                    # Go module definition
└── graph/
    ├── schema.graphqls       # GraphQL schema
    ├── resolver.go           # Resolver struct and seed data
    ├── schema.resolvers.go   # Resolver implementations
    ├── generated.go          # Auto-generated (do not edit)
    └── model/
        └── models_gen.go     # Auto-generated models (do not edit)
```
