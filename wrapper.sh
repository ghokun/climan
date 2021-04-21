#!/bin/bash

# Check if it is installed
climan install $0
echo $0 is called

# Run original command after installation
$0 "$@"
