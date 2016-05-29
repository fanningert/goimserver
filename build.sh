#!/bin/bash

export PROJECT_ROOT=$(pwd)

# Test
rm glide.lock
rm glide.yaml

# Prepare GO build enviroment
export GOROOT=/usr/lib/go

rm -rf build
mkdir -p $PROJECT_ROOT/build/go
cd $PROJECT_ROOT/build/go

for f in "$GOROOT/"*; do
  ln -s "$f"
done

rm pkg
mkdir pkg
cd pkg

for f in "$GOROOT/pkg/"*; do
  ln -s "$f"
done

export GOROOT="$PROJECT_ROOT/build/go"
export GOPATH="$PROJECT_ROOT/build"

cd $GOPATH
rm src
mkdir src

cd $PROJECT_ROOT
glide create
glide install

cd $PROJECT_ROOT/build/src
for f in "$PROJECT_ROOT/vendor/"*; do
  ln -s "$f"
done

cd $PROJECT_ROOT
go build
#glide up