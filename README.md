# grpc_hello_world
GRPC Hello World

#Install grpc
```
go get -u google.golang.org/grpc
```

#Install the protoc Go plugin
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#To generate the GRPC Proto File
```
protoc -I . helloworld.proto --go_out=plugins=grpc:.
```
