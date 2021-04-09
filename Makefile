all: build

build:
	go build -o dist/nset-cli

install:
	\cp dist/nset-cli /opt/nset-ansible/
