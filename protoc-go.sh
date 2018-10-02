#!/bin/bash -e

#PROTOC_FILE="protoc-3.6.1-linux-x86_64.zip"
PROTOC_FILE="protoc-3.5.1-linux-x86_64.zip"
PROTOC_PATH="protobuf/protoc"
PROTOC="protobuf/protoc/bin/protoc"

OS=$(uname -s)
if [ $OS = "Darwin" ]; then
	#PROTOC_FILE="protoc-3.6.1-osx-x86_64.zip"
	PROTOC_FILE="protoc-3.5.1-osx-x86_64.zip"
fi

# setup for the first time
#go get github.com/favadi/protoc-go-inject-tag
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
rm -fr $PROTOC_PATH && mkdir -p $PROTOC_PATH
unzip -q protobuf/$PROTOC_FILE -d $PROTOC_PATH

if [ ! -f $GOPATH/bin/protoc-gen-go ]; then
	echo -e "\033[0;34minstalling protoc-gen-go..."
	go get -u github.com/golang/protobuf/protoc-gen-go
	echo "done"
fi

if [ ! -f $GOPATH/bin/protoc-gen-swagger ]; then
	echo -e "\033[0;34minstalling proto-gen-swagger..."
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	echo "done"
fi

[ "$1" = 'all' ] && if [ ! -f ./node_modules/protobufjs/bin/pbjs ]; then
	echo -e "\033[0;34minstalling protobufjs..."
	npm i
	echo "done"
fi

TOTAL=1
echo -e "\033[0;34mcompiling proto files"
$PROTOC --version

mkdir -p scala
[ "$1" = 'all' ] && rm scala/* -Rf
ALLPROTO=""

for i in `ls -R`; do
	if [[ $i == *":"* ]]; then
		LAST_DIR=${i%?} # trim last char
		continue
	fi

	if [[ $i == "vendor" ]] || [[ $i == "proto" ]] || [[ $LAST_DIR == ./node_modules* ]] || [[ $LAST_DIR == ./vendor* ]] || [[ $LAST_DIR == ./proto* ]]; then
		continue
	fi

	if [[ $i == *".proto" ]]; then
		echo -e "\033[0;90m["$TOTAL"] compiling" $LAST_DIR /$i "\033[0;31m" &
		$PROTOC --go_out=plugins:. --proto_path=$GOPATH/src --proto_path=./  $LAST_DIR/$i &
		$PROTOC -I$PROTOC_PATH/include -I. -I$GOPATH/src --swagger_out=logtostderr=true:. --proto_path=$GOPATH/src --proto_path=./ $LAST_DIR/$i
		#./protoc --python_out=plugins:. --proto_path=$GOPATH/src --proto_path=./ $LAST_DIR/$i
		# ./protoc -I=$GOPATH/src --java_out=. $GOPATH/src/bitbucket.org/subiz/header/$LAST_DIR/$i
		ALLPROTO="$ALLPROTO $LAST_DIR/$i"
		let "TOTAL += 1"
		# else
		# echo -e "\033[0;37mignore" $i "\033[0;30m"
		[ "$1" = 'all' ] && ./bin/scalapbc -v351 -I $GOPATH/src --scala_out=flat_package:./scala $GOPATH/src/git.subiz.net/header/$LAST_DIR/$i > /dev/null
	fi
done;

wait

[ "$1" = 'all' ] && echo -e "\033[0;90m["$TOTAL"] generating subiz.d.ts\033[0;31m"
[ "$2" = 'all' ] && npx pbjs --es6 --keep-case --no-convert --no-create --no-encode --no-decode --no-verify --no-delimited --no-beautify -t static-module -w commonjs -o subiz.js $ALLPROTO
[ "$1" = 'all' ] && npx pbts --no-comments -o subiz.d.ts subiz.js

rm -fr $PROTOC_PATH

echo -e "\033[0;32mDone. (total:" $TOTAL "files)\033[0;30m"
