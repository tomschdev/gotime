# go-concurrency

This repo holds implementations of thread-safe go servers presented in the LinkedIn course [Applied Concurrency in Go](https://www.linkedin.com/learning/applied-concurrency-in-go). \
The intention is to use this repo as a reference for best practices in API concurrency.
## sync-server
Enforces concurrency-safe hanlders using the sync package, provided by go.
## chan-server
Improves on sync-server implementation by using built in locks of go channels.
## Execution
The execution of all implementations follows the same procedure, chan-server will be used as an example:
```bash
cd chan-server
go run server.go
```
From new terminal session:
```bash
cd chan-server
go run cmd/simulate.go
```


