#!/bin/bash
if [ -e foobar ]; then
  rm -rf foobar
fi

go run cmd/zona/main.go test

bat foobar/in.html
