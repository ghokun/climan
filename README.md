# climan

`climan` is a command line tool version manager for cloud native technologies.

## Features
- No root user is required.
- Switch between versions.
- All binaries are linked by default. They are installed upon first call.
- Sync client with api version.

## Installation
### Automated
```bash
# Fancy shell script
```

### Manual
```bash
# Download binary from releases page

# Make it executable
chmod +x climan
./climan install climan

# Add climan bin folder to your PATH
export PATH="$HOME/.climan/bin:$PATH"
```

## Usage
```bash
# Help
climan help

# List all tools (hit cache first, then remote)
climan list

# List all tools (only from remote)
climan list --force-remote

# List all  versions of a tool
climan list kubectl

# Install the latest version of a tool
# Sets the latest version as default
climan install kubectl

# Install the latest version of a tool but don't set as default
climan install kubectl --no-default

# Install a specific version of a tool
climan install kubectl 1.20

# Remove the latest version of a tool
climan remove kubectl

# Remove a specific version of a tool
climan remove kubectl 1.20

# Remove all versions of a tool
climan remove kubectl --all
```

## Tools
```bash
climan # Yes, you can download any version of this tool to manage itself

# API Clients
kubectl
oc
kn
kamel
odo

# Management
helm
kustomize

# Pipelines
flux
tkn
argocd

# Distributions
minikube
kind
k3d
crc
```
## Development
This repository contains `devcontainer` files that can be opened in Visual Studio Remote Containers.

## TASKS
