# Rest api
1. clone swagger-ui
2. access http://src/swagger-ui/dist/index.html
3. paste http://src/servicespec/rest/account.yaml to input box

# How to build
## build proto
```
./protoc-go.sh
```
## build rest
1. npm run rest

## build all
1. ./build.sh

# How to update protoc

Download newer version from https://github.com/google/protobuf/releases
the current version is `protobuf-cpp-3.5.1.tar.gz`
```
tar -xzvf protobuf-cpp-3.5.1.tar.gz
cd protobuf-cpp-3.5.1
./autogen.sh

./configure --disable-shared
make -j7
make check -j7
sudo make install
# sudo ldconfig # refresh shared library cache.
# a new protoc will be located at
which protoc

copy it to src/bitbucket.org/subiz/header/protobuf/bin
# zip it with
tar -zcvf protoc
rm protoc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

For any error, please consume README.md in `protobuf-cpp-3.5.1` or `protobuf-cpp-3.5.1/src`
