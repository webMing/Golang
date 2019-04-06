#!/bin/bash
clear 
export GOPATH="$(go env GOPATH)":"$PWD"
#export GOPATH="go env GOPATH":"$PWD"
echo "Current GOPATH:$GOPATH"
echo "Para:$1"
go run $1
