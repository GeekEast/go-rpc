## Install
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

## Protocol Buffers
- the default method of serializing structured data in `gRPC`
### What for
- serializing structured data
### Advantages
- faster for serialzing than JSON, XML
- faster for transmiting than JSON, XML

### Use Case
- transmit data between microservices

### Understanding
- json
```json
{
  first_name: "Arun",
}
```
- proto
```proto
{
  string first_name = 1;
}
```
- data in protobuf message
```
124Arun
```
> 1: first_name (field name)
> 2: string (field type)
> 4: 4 bytes (field size)

### Format
- field rule: `repeated` to represent an **array** of same type data
```proto
message FileDescriptorProto {
    // Names of files imported by this file.
    repeated string dependency = 3;
}
```
- field type: [scalar](https://developers.google.com/protocol-buffers/docs/proto3#scalar) and [enum](https://developers.google.com/protocol-buffers/docs/proto3#enum)
```proto
message Person {
  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }
}
```
- field name: in **lowercase**, separated by `_`
```proto
message Person {
  bool is_active = 3;
}
```
- field tag: **unique**, **integer**; [reserved](https://developers.google.com/protocol-buffers/docs/proto3#reserved)
```proto
message PhoneNumber {
    string number = 1;
}

message Person {
    uint64 id = 1;
    string email = 2;
}
```
## Proto3
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

## Practice
- no `*` and `&` for interface
- `*` and `&` for `struct` (by now)

## References
- [gRPC](https://grpc.io/docs/guides/)
- [gRPC to Typescript](https://github.com/improbable-eng/ts-protoc-gen)
- [Understand Protocol Buffer](https://medium.com/better-programming/understanding-protocol-buffers-43c5bced0d47)
- [Google Protocol Buffer 3](https://developers.google.com/protocol-buffers/docs/proto3#simple)
- [Protocol Buffer vs JSON](https://auth0.com/blog/beating-json-performance-with-protobuf/)