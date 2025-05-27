#!/bin/bash

# Go to project root
cd ../src
# Start backend in background
go run -C backend_go . &

# Start frontend in foreground
python3 -m http.server 2020
