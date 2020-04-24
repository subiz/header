#!/bin/bash -e

TOTAL=1
echo -e "\033[0;34mcompiling proto files"
$PROTOC --version

ALLPROTO=""
cd /src

for i in `ls -R`; do
	if [[ $i == *":"* ]]; then
		LAST_DIR=${i%?} # trim last char
		continue
	fi

	# ignore some dirs
	if [[ $i == "vendor" ]] || [[ $i == "proto" ]] || [[ $LAST_DIR == ./node_modules* ]] || [[ $LAST_DIR == ./vendor* ]] || [[ $LAST_DIR == ./proto* ]]; then
		continue
	fi

	if [[ $i == *".proto" ]]; then
		printf "\033[0;90m["%d"] compiling %s %s \033[0;31m\n" $TOTAL $LAST_DIR /$i
		$PROTOC -I/tmp/include -I. --go_out=plugins=grpc:. --proto_path=./  $LAST_DIR/$i &

		ALLPROTO="$ALLPROTO $LAST_DIR/$i"
		let "TOTAL += 1"
	fi
done;
wait

cp -r ./github.com/subiz/header/* ./
rm -rf github.com

echo -e "\033[0;32mDone. (total:" $TOTAL "files)\033[0;30m"
