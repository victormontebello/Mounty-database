#!/bin/bash

# Check for open ports
echo "Checking for open ports..."
open_ports=$(netstat -tuln)
if [[ -n "$open_ports" ]]; then
    echo "Warning: Open ports detected!"
    echo "$open_ports"
else
    echo "No open ports detected."
fi

# Check file permissions
echo "Checking file permissions..."
file_permissions=$(ls -l data.json)
if [[ "$file_permissions" =~ "-rw-r--r--" ]]; then
    echo "File permissions OK."
else
    echo "Warning: Insecure file permissions detected!"
    echo "$file_permissions"
fi
