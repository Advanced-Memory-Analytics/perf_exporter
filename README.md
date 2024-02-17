## Build the Container
`docker build -t apramay/docker-perf-test:latest -f Dockerfile .`

## Run the Container:
`docker run --rm --name apramay_perf_scraper --security-opt seccomp=seccomp-perf.json --privileged -d apramay/docker-perf-test:latest`