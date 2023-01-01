# Go Elasticsearch

## Projects Structure

```
<root>
    ├── Makefile
    ├── cmd
    │   └── elastic
    │       └── main.go
    ├── docker
    │   ├── docker-compose.yaml
    │   └── tmp
    │       └── elasticsearch
    └── internal
        ├── pkg
        │   ├── domain
        │   │   └── error.go
        │   └── storage
        │       ├── elasticsearch
        │       │   ├── elasticsearch.go
        │       │   └── post_storage.go
        │       └── post_storer.go
        └── post
            ├── handler.go
            ├── request.go
            ├── response.go
            └── service.go
```

## Prerequisites

1. Golang (v1.16.4 or above)
2. Elasticsearch

## Installation

1. git clone git@github.com:geekbim/Go-Elasticsearch.git
2. sync go modules -> `go mod tidy`
3. download required modules -> `go mod download`
4. testing
```
    curl --request POST 'http://localhost:3000/api/v1/posts' \
    --data-raw '{
        "title": "test title",
        "text": "test text",
        "tags": ["tag"]
    }'
     
    curl --request PATCH 'http://localhost:3000/api/v1/posts/{id}' \
    --data-raw '{
        "title": "test title x",
        "text": "test text x",
        "tags": ["tag x"]
    }'
     
    curl --request DELETE 'http://localhost:3000/api/v1/posts/{id}'
     
    curl --request GET 'http://localhost:3000/api/v1/posts/{id}'
```