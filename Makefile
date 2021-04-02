all: build install

build:
	go build -o nset-cli

install:
	\cp launch /opt/nset-ansible/
