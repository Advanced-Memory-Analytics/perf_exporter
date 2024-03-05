#!/bin/bash

echo "Script execution starting..." >> /log.txt

# Set kernel.perf_event_paranoid to -1
echo "-1" > /proc/sys/kernel/perf_event_paranoid

# Running exporter
(cd /app/ &&  go run cmd/main.go)

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?
