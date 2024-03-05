.PHONY: build run all

build:
	(cd exporter/ && go mod tidy)

run: 
	(cd exporter/cmd && go run main.go)


all: 
	(cd exporter/ && go mod tidy)
	(cd exporter && go run cmd/main.go)