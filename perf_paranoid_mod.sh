#!/bin/bash

# Set kernel.perf_event_paranoid to -1
echo "-1" > /proc/sys/kernel/perf_event_paranoid

# Executing the command passed to the entrypoint
exec "$@"