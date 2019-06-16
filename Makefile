#!/usr/bin/env bash

fmt:
	gofmt -l -w ./

install:fmt
	go build -v -o output/fileServe ./
