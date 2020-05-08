### Install
- **Golang** `go mod init go-rpc`
```sh
# ~/.zshrc
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/Cellar/go/1.14.2_1/bin
```
- **Vscode Extensions**
    - **vscode-proto3**
    - **Clang-Format** `brew install clang-format`
- **protoc** `brew install protobuf`
- **go-gRPC**
```sh
go get -u google.golang.org/grpc
```
- **protoc-gen-go**
```sh
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Proto3
> use this to generate gRPC `service` code for both **server-side** and **client-side** based on protocol (not the server code!)
- First `.proto`
```proto3
syntax = "proto3";
package greet;
option go_package="greetpb";

service GreetService {}
```
- Use `.proto` to generate go code
```sh
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```