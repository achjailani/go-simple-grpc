# GO Simple gRPC

## Installation
Clone repository

```shell
git clone https://github.com/achjailani/go-simple-grpc.git
```
Install

```shell
cd go-simple-grpc && go mod download 
```

## Running Application
Run gRPC server
```shell
go run main.go grpc:start
```
Run HTTP Server
```shell
go run main.go
```

## gRPC Server
Any related things to gRPC server could be found in `grpc` directory, including `handler`, `server` and `interceptors`

## gRPC Client
Any related things to gRPC client could be found in `client` directory, including `connection`, `method call`

## REST API
Any related things to REST API could be found in `rest` directory