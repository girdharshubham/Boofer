![Go](https://github.com/girdharshubham/Boofer/workflows/Go/badge.svg?branch=master)
# Boofer
Gin is a web framework written in Go (Golang). This template provides a ready to use template for building web applications. It includes a router group(**/v1**) set up with basic auth, 67% unit test coverage right out of the box, a basic build using github actions, and a makefile to simplify building this app.

## Directory Structure
```
.
├── README.md
├── api
│   └── v1.go
├── bin
│   └── boofer
├── go.mod
├── main.go
├── makefile
└── service
    ├── allservices.go
    └── allservices_test.go

```

## Routes
 ```
 GET /v1/healthz
 GET /v1/about
 GET /v1
 ```

## How to run
**Prerequisite:** You should have Go set up on your system.
 * Run test cases with coverage
   ```
   make test
   ```
 * View Coverage
   ```
   make view
   ```
 * Building the binairy
   ```
   make package
   ```
 * Run the project
   ```
   make run
   ```
 * Clean up
   ```
   make clean
   ```
