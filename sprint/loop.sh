#!/bin/bash

# Check if the number of arguments is correct
if [ $# -ne 1 ]; then
    echo "Usage: $0 <number>"
    exit 1
fi

# Get the first argument
count=$1

# Limit the loop to 100 times if the argument is greater than 100
if [ $count -gt 100 ]; then
    count=100
fi

# Loop and print the message
for (( i=1; i<=count; i++ ))
do
    echo "$i times I've printed hannesverlis1"
done