#!/bin/bash
clear 
export GOPATH="$PWD"
echo "Current GOPATH:$GOPATH"
echo "Para:$1"
go run $1
