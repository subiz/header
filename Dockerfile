FROM golang:1.13

WORKDIR /tmp

RUN apt update && apt install unzip

# install protobuf v3.11.2
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip -O protoc.zip
RUN unzip protoc.zip

# install protoc gengo
RUN GO111MODULE=on go get -u github.com/golang/protobuf/protoc-gen-go@v1.4.0
RUN GO111MODULE=on go get google.golang.org/grpc@v1.28.1

ENV PROTOC=/tmp/bin/protoc

COPY protoc-go.sh ./
ENTRYPOINT /tmp/protoc-go.sh
