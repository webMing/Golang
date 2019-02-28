#!/bin/bash
clear 
export GOPATH="$PWD"
echo "Current GOPATH:$GOPATH"
go run $1
