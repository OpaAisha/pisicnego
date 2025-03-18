#!/bin/bash

# Use find to count regular files and directories
count=$(find . -type f -o -type d | wc -l)

# Print the count
echo $count
