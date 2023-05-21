obj=corednswebui
installPath=/usr/bin

all:build

install:
	go mod tidy
	go build -o ${installPath}/${obj} main.go

build:
	go mod tidy
	go build -o ${obj} main.go