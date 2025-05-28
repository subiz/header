FROM golang:1.24.2
WORKDIR /tmp
RUN apt update && apt install -y unzip curl

# install node
ENV NODE_VERSION=22.16.0
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.40.3/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

# install protobuf v28.3
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v31.0/protoc-31.0-linux-x86_64.zip -O protoc.zip
RUN unzip protoc.zip

# install protoc gengo
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

ENV PROTOC=/tmp/bin/protoc

RUN echo {} > package.json
RUN npm i --save lodash
RUN npm i --save @subiz/langmap
RUN cat ./package.json
COPY lang.js ./
COPY protoc-go.sh ./

ENTRYPOINT ["/tmp/protoc-go.sh"]
