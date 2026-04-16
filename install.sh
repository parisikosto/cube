#!/bin/bash

set -e

LATEST_RELEASE="https://github.com/parisikosto/cube/releases/latest/download/cube-linux-amd64"

echo "Downloading cube..."
curl -fsSL -o cube $LATEST_RELEASE

chmod +x cube
sudo mv cube /usr/local/bin/cube

echo "Setting up shell completion..."
if [ -d /etc/bash_completion.d ]; then
  cube completion bash | sudo tee /etc/bash_completion.d/cube > /dev/null
fi

echo "Installation complete! Run 'cube --help' to get started."
