#!/usr/bin/env bash

v=$(cat VERSION)

echo "Building: stefanop/home-hap:$v"

go build \
    -ldflags="-X 'home-hap/internal/configs.Version=$version'" \
    -o ./dist/home-hap ./cmd/server