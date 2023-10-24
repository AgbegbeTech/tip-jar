#!/bin/bash

# Build the Go application
go build main.go

# Ensure the config.json file is present
if [ ! -f config.json ]; then
    echo "Error: config.json file is missing. Create a valid configuration file."
    exit 1
fi

# Run the application with nohup (you can replace 'main' with the actual binary name)
nohup ./main > app.log 2>&1 &

# Optionally, you can perform other deployment tasks here.
# For example, configuring reverse proxies, setting up a web server, etc.

echo "Application deployed successfully."
