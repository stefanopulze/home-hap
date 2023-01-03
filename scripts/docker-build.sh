#!/usr/bin/env bash

v=$(cat VERSION)
r=registry.home.stefanopulze.dev

echo "Building docker image: home-hap:$v"

docker build -t $r/home-hap:$v --build-arg version=$v .
#docker push $r/home-hap:$v