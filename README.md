## Build the Container
`docker build -t apramay/docker-perf-test:latest -f Dockerfile .`

## Run the Container:
`docker run --rm --name apramay_perf_scraper -p 9100:9100 --security-opt seccomp=seccomp-perf.json --privileged -d apramay/docker-perf-test:latest`
