#!/bin/bash

# Get the latest version
VERSION=$(curl -s https://api.github.com/repos/TechSquidTV/uhs-cli/releases/latest | grep tag_name | cut -d '"' -f 4)
RED='\033[0;31m'
NC='\033[0m' # No Color

# The operating system and architecture
OS=$(uname)
ARCH=$(uname -m)

# The URL of the release
URL="https://github.com/TechSquidTV/uhs-cli/releases/download/${VERSION}/uhs-cli_${OS}_${ARCH}.tar.gz"

# Download and extract the binary
wget $URL
mkdir uhs-cli_${OS}_${ARCH}
tar -xvzf uhs-cli_${OS}_${ARCH}.tar.gz -C uhs-cli_${OS}_${ARCH}

# Move the binary to a location in your PATH
if [ ! -d ~/.local/bin ]; then
  mkdir -p ~/.local/bin
fi
mv uhs-cli_${OS}_${ARCH}/uhs-cli ~/.local/bin/uhs

# Ensure the binary is executable
chmod +x ~/.local/bin/uhs

# Clean up the downloaded file and the extracted directory
rm uhs-cli_${OS}_${ARCH}.tar.gz
rm -rf uhs-cli_${OS}_${ARCH}

echo "Installed uhs-cli version ${VERSION}"
echo "Run 'uhs --help' to get started"
echo 
# Check if the binary is in your PATH
if ! command -v uhs &> /dev/null; then
  echo "${RED}[WARNING]${NC} uhs could not be found in your PATH. Please add '\${HOME}/.local/bin/' to your PATH."
fi
