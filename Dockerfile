FROM golang:1.17

WORKDIR /tmp

RUN apt update && apt install unzip

# install node 12.6.0
ENV NODE_VERSION=12.6.0
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.34.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

# install protobuf v21.3
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v21.3/protoc-21.3-linux-x86_64.zip -O protoc.zip
RUN unzip protoc.zip

# install protoc gengo
RUN go get -u github.com/golang/protobuf/protoc-gen-go
# RUN go get -v google.golang.org/grpc

ENV PROTOC=/tmp/bin/protoc

RUN echo {} > package.json
RUN npm i --save lodash
RUN npm i --save langmap
RUN ls
COPY lang.js ./
COPY protoc-go.sh ./

ENTRYPOINT /tmp/protoc-go.sh
