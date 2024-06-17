#!/bin/bash

cd ../internal/protobufs

protoc *.proto --go_out=.. --go-grpc_out=..

cd ../services

for dir in */ ; do
    # Check if the item is a directory
    if [ -d "$dir" ]; then
        echo "Entering directory: $dir"
        cd "$dir" || exit
        # Execute the specified command
        protoc *.proto --go_out=.. --go-grpc_out=..
        # Return to the parent directory
        cd ..
    fi
done