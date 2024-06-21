FROM golang:1.22.1

WORKDIR /tmp

RUN apt update && apt install -y unzip curl

# install node 20.15.0
ENV NODE_VERSION=20.15.0
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.39.7/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

# install protobuf v21.12
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v27.1/protoc-27.1-linux-x86_64.zip -O protoc.zip
RUN unzip protoc.zip

# install protoc gengo
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.4

ENV PROTOC=/tmp/bin/protoc

RUN echo {} > package.json
RUN npm i --save lodash
RUN npm i --save langmap
RUN ls
COPY lang.js ./
COPY protoc-go.sh ./

ENTRYPOINT /tmp/protoc-go.sh
