#!/usr/bin/env bash

fmt:
	gofmt -l -w ./

dev:fmt
	go build -v -o output/fileServe ./
