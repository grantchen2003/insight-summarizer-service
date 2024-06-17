#!/bin/bash

cd ..

export ENV=dev

nodemon --exec "go run cmd/summarizer/main.go" --watch . -e go