#!/bin/bash

# Go to project root
cd ..
source venv/bin/activate
# Start backend in background
python3 src/backend/main.py &

# Start frontend in foreground
cd src
python3 -m http.server 2020
