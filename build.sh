docker build -t live360vn/protobuild:2.3 .
docker run -it --rm -v `pwd`:/src live360vn/protobuild:2.3
