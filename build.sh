#!/bin/bash

echo "Script execution starting..." >> /log.txt
# Set kernel.perf_event_paranoid to -1
echo "-1" > /proc/sys/kernel/perf_event_paranoid

# Running Perf_Exporter
(cd /app/ &&  go run cmd/main.go)
# ./app/main


# Running Node_Exporter
# (cd ../../ && ./usr/local/bin/node_exporter) &
cd ../ && ./usr/local/bin/node_exporter

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?

# # Executing the command passed to the entrypoint
# exec "$@"