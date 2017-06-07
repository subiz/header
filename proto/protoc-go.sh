TOTAL=0
echo -e "\033[0;34m===========BUILDING PROTO SPEC======="

for i in `ls -R`; do
  if [[ $i == *":"* ]]; then
    LASTDIRECTORY=${i::-1}
  else
    if [[ $i == *".proto" ]]; then
      echo -e "\033[0;37mcompiling" $LASTDIRECTORY/$i "\033[0;31m"
      protoc --go_out=plugins:. --proto_path=../../../../ --proto_path=./  $LASTDIRECTORY/$i
			protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. --proto_path=../../../../ --proto_path=./ $LASTDIRECTORY/$i
      let "TOTAL += 1"
    # else
     # echo -e "\033[0;37mignore" $i "\033[0;30m"
    fi;
  fi;
done;
echo -e "\033[0;32mDone. (total:" $TOTAL "files)\033[0;30m"

