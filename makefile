#!/bin/sh

workdir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

wd:
	echo $(workdir)

build:
	go build -C ./backend/cmd/api -o ../../bin/api

build-fs:
	go build -C ./backend/cmd/serve -o ../../bin/serve

run:
	# sudo lsof -i :5000
	# kill PID
	./bin/api &
