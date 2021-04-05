# climan

`climan` is a command line tool version manager for cloud native technologies.

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
# List all tools
climan list

# List all  versions of a tool
climan list kubectl

# Install the latest version of a tool
climan install kubectl

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
