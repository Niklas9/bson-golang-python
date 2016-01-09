#!/bin/sh

# Install needed BSON dependencies, both for Go and Python

export GOPATH=$(dirname $(readlink -f "$0"))  # this is cwd

# get packages
go get gopkg.in/mgo.v2/bson 
sudo pip install -r requirements.txt
