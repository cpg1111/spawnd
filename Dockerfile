FROM golang:1.7
RUN apt-get update && \
    apt-get install -y curl build-essential && \
    curl https://glide.sh/get | sh
COPY . $GOPATH/src/github.com/cpg1111/spawnd/
WORKDIR $GOPATH/src/github.com/cpg1111/spawnd
ENTRYPOINT ["bash"]
