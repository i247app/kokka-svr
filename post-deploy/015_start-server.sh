#!/usr/bin/env bash

KOKKA_HOME="/apps/kokka"
cd $KOKKA_HOME

echo "$KOKKA_HOME"
echo "Starting new server..."

# Store output in a logfile and save the PID to a file so we can kill the process later
./dist/server >> /apps/kokka/gosvr.log 2>&1 & echo $! > /apps/kokka/gosvr.pid

echo "Verifying the server is running..."
if ! ps -p $(cat /apps/kokka/gosvr.pid) > /dev/null 2>&1; then
    echo "ERROR: Process is not running!"
    exit 1
fi
echo "OK!"
