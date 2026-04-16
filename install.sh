#!/bin/bash

set -e

# Define the latest release URL
LATEST_RELEASE="https://github.com/parisikosto/cube/releases/latest/download/cube-linux-amd64"

# Download the binary
echo "Downloading cube-linux-amd64..."
curl -L -o cube $LATEST_RELEASE

# Move it to /usr/local/bin
chmod +x cube
sudo mv cube /usr/local/bin/cube

echo "Installation complete! Run 'cube --help' to get started."
