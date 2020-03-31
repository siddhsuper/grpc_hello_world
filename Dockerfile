FROM golang:1.14
RUN mkdir -p /tmp/src/github.com
WORKDIR /tmp/src/github.com/grpc_hello_world/
ADD . /tmp/src/github.com/grpc_hello_world
RUN GOPATH=/tmp && \
    cd /tmp/src/github.com/grpc_hello_world && \
    go get &&\
    go build
EXPOSE 8888
CMD ["./grpc_hello_world"]

