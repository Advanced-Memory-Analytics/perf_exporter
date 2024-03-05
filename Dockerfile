FROM debian:bullseye-slim

ARG ARCH_TYPE="arm64"
ENV ARCH_TYPE=$ARCH_TYPE

COPY --from=golang:1.22-bullseye /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

# Install dependencies
RUN apt-get update \
    && apt-get install -y \
        linux-perf \
        linux-base \
        net-tools

RUN rm -f /usr/bin/perf
RUN mv /usr/bin/perf_5.10 /usr/bin/perf

# Download and install node_exporter
RUN apt-get update && apt-get install -y procps
RUN apt-get install -y ca-certificates wget build-essential
RUN wget https://github.com/prometheus/node_exporter/releases/download/v1.7.0/node_exporter-1.7.0.linux-${ARCH_TYPE}.tar.gz
RUN tar xvfz node_exporter-1.7.0.linux-${ARCH_TYPE}.tar.gz
RUN cp node_exporter-1.7.0.linux-${ARCH_TYPE}/node_exporter /usr/local/bin/

# Cleanup
RUN rm -rf node_exporter-1.7.0.linux-${ARCH_TYPE}/ node_exporter-1.7.0.linux-${ARCH_TYPE}.tar.gz

# Expose the default port used by node_exporter
EXPOSE 9100
EXPOSE 8080

COPY docker-build.sh /usr/local/bin/
COPY exporter /app

# Set permissions and working directory
RUN chmod +x /usr/local/bin/docker-build.sh
RUN chmod +x /app/main

# Run node_exporter
RUN ls -l /usr/local/bin/docker-build.sh
CMD ./usr/local/bin/docker-build.sh
