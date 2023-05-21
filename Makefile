obj=corednswebui

all:build

build:
	go mod tidy
	go build -o ${obj} main.go