all: build

get-deps:
	cd ${GOPATH} && \
	go get -u github.com/kardianos/govendor

build:
	govendor sync
	go build -o spawnd-container container/container.go
	go build -o spawnd main.go

test:
	go test -v ./...

install:
	mkdir -p /etc/spawnd/conf.toml
	mkdir -p /usr/bin/spawn.d/
	cp test_conf.toml /etc/spawnd/conf.toml
	cp spawnd-container /usr/bin/spawn.d/
	cp spawnd /usr/bin/spawn.d/
	echo "#!/bin/bash\n/usr/bin/spawn.d/spawnd $@\n" > /usr/bin/spawnd
	chmod 555 /usr/bin/spawnd
	
