# Simple gRPC

## Installation

For first installation we need to install proto compiler in our machine:

- Mac: ` brew install protobuf `
- Linux: ` apt install -y protobuf-compiler `

Update your PATH so that the protoc compiler can find the plugins:

- ` export PATH="$PATH:$(go env GOPATH)/bin" `

Initialize go module:

- ` go mod init github.com/irhamsahbana/simple-grpc `

Install the protocol compiler plugins for Go using the following commands:

- ` go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 `
- ` go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 `

## Compile proto file

For this project we can use command below for compile proto file inside student folder:

```shell
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. student/student.proto
```

## Run gRPC server

We need jump into `server` folder and then run `main.go` file

```shell
cd server && go run main.go
```

## Run gRPC client

We need change directory to `client` folder and run the entry point

```shell
cd client && go run main.go
```

After we run gRPC client entrypoint we will triggered the request function and trigger gRPC server.
