FROM debian:bullseye-slim

ARG ARCH_TYPE="arm64"

ENV ARCH_TYPE=$ARCH_TYPE

# Install dependencies
# RUN apt-get update && apt-get install -y wget
RUN apt-get update \
    && apt-get install -y \
        linux-perf \
        linux-base \
        net-tools

# # Set up Perf Permissions
# RUN su -c 'echo -1 > /proc/sys/kernel/perf_event_paranoid' -s /bin/sh root

# Download and install node_exporter
RUN apt-get install -y ca-certificates wget build-essential
RUN wget https://github.com/prometheus/node_exporter/releases/download/v1.7.0/node_exporter-1.7.0.linux-${ARCH_TYPE}.tar.gz
RUN tar xvfz node_exporter-1.7.0.linux-${ARCH_TYPE}.tar.gz
RUN cp node_exporter-1.7.0.linux-${ARCH_TYPE}/node_exporter /usr/local/bin/

# Cleanup
RUN rm -rf node_exporter-1.7.0.linux-armm64/ node_exporter-1.7.0.linux-arm64.tar.gz


# Expose the default port used by node_exporter
EXPOSE 9100

# Run node_exporter
ENTRYPOINT ["/usr/local/bin/node_exporter"]