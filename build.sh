#!/bin/bash

go mod tidy
go build -o bin/zona ./cmd/zona
sudo cp -f bin/zona /usr/bin/zona
