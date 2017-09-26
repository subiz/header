# setup
#go get github.com/favadi/protoc-go-inject-tag

TOTAL=1
echo -e "\033[0;34mcompiling proto files"

for i in `ls -R`; do
  if [[ $i == *":"* ]]; then
    LASTDIRECTORY=${i::-1}
  else
    if [[ $i == *".proto" ]]; then
      echo -e "\033[0;90m["$TOTAL"] compiling" $LASTDIRECTORY/$i "\033[0;31m"
      protoc --go_out=plugins:. --proto_path=../../../ --proto_path=./  $LASTDIRECTORY/$i
			protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. --proto_path=../../../ --proto_path=./ $LASTDIRECTORY/$i
			protoc --python_out=plugins:. --proto_path=../../../ --proto_path=./ $LASTDIRECTORY/$i
      let "TOTAL += 1"
    # else
     # echo -e "\033[0;37mignore" $i "\033[0;30m"
    fi;
  fi;
done;
echo -e "\033[0;34mremoving all omitempty\033[0;31m"
# ls */*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
let "TOTAL -= 1"

#echo "\033[0;34minjecting golang custom tag\033[0;31m"
#ls */*.pb.go | xargs -n1 -IX bash -c 'protoc-go-inject-tag -input=X'
echo -e "\033[0;32mDone. (total:" $TOTAL "files)\033[0;30m"
