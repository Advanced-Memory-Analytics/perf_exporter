.PHONY: build run all

build:
	(cd perf_exporter/ && go mod tidy && go mod vendor)
	chmod +x perf_exporter/load.sh

run: 
	(cd perf_exporter/cmd && go run main.go)


all: 
	build 
	run