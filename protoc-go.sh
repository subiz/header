# setup for the first time
#go get github.com/favadi/protoc-go-inject-tag
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
tar xzvf protobuf/bin/protoc.tar.gz

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

TOTAL=1
echo -e "\033[0;34mcompiling proto files"

for i in `ls -R`; do
		if [[ $i == *":"* ]]; then
				LASTDIRECTORY=${i::-1}
		else
				if [[ $i == "vendor" ]] || [[ $i == "proto" ]] || [[ $LASTDIRECTORY == ./vendor* ]] || [[ $LASTDIRECTORY == ./proto* ]]; then
						continue
				fi
				if [[ $i == *".proto" ]]; then
						echo -e "\033[0;90m["$TOTAL"] compiling" $LASTDIRECTORY /$i "\033[0;31m"
						./protoc --go_out=plugins:. --proto_path=$GOPATH/src --proto_path=./  $LASTDIRECTORY/$i
						./protoc -I./protobuf/include -I. -I$GOPATH/src --swagger_out=logtostderr=true:. --proto_path=$GOPATH/src --proto_path=./ $LASTDIRECTORY/$i
						#./protoc --python_out=plugins:. --proto_path=$GOPATH/src --proto_path=./ $LASTDIRECTORY/$i
						let "TOTAL += 1"
						# else
						# echo -e "\033[0;37mignore" $i "\033[0;30m"
				fi;
		fi;
done;
# echo -e "\033[0;34mremoving all omitempty\033[0;31m"
# ls */*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
let "TOTAL -= 1"

#echo "\033[0;34minjecting golang custom tag\033[0;31m"
#ls */*.pb.go | xargs -n1 -IX bash -c 'protoc-go-inject-tag -input=X'
rm protoc
echo -e "\033[0;32mDone. (total:" $TOTAL "files)\033[0;30m"
