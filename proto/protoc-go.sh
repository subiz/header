TOTAL=0
for i in `ls -R`; do
  if [[ $i == *":"* ]]; then
    LASTDIRECTORY=${i::-1}
  else
    if [[ $i == *".proto" ]]; then
      echo -e "\033[0;37mcompiling" $LASTDIRECTORY/$i "\033[0;31m"
      protoc --go_out=plugins:. --proto_path=../../../../ --proto_path=./  $LASTDIRECTORY/$i
      let "TOTAL += 1"
    # else
     # echo -e "\033[0;37mignore" $i "\033[0;30m"
    fi;
  fi;
done;
echo -e "\033[0;34mDone. (total:" $TOTAL "files)\033[0;30m"

