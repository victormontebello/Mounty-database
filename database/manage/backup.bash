#!/bin/bash

# Backup directory
backup_dir="/path/to/backup/directory"

source_file="data.json"

destination_file="$backup_dir/data_$(date +"%Y%m%d%H%M%S").json"

if [ ! -d "$backup_dir" ]; then
    mkdir -p "$backup_dir"
fi

if [ ! -f "$source_file" ]; then
    echo "Error: Source file '$source_file' not found."
    exit 1
fi

cp "$source_file" "$destination_file"

if [ $? -eq 0 ]; then
    echo "Backup created successfully: $destination_file"
else
    echo "Error creating backup."
fi
