#!/bin/bash

# Help message
help_message="
Usage: ./help.sh <command>

Commands:
  insert <key> <value>   Insert a new record with the specified key and value.
  delete <key>            Delete the record with the specified key.
  update <key> <value>   Update the value of an existing record with the specified key.
  select <key>            Select and print the record with the specified key.
  selectall               Select and print all records.
  help                    Show this help message.
"

# Print help message
echo "$help_message"
