# docker build -t live360vn/protobuild:2.9 .
# docker build -t asia-southeast1-docker.pkg.dev/subiz-version-4/subiz/test:1 .
docker run -it --rm -v `pwd`:/src live360vn/protobuild:2.9
