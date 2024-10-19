#!/bin/bash

go build -o bin/zona ./cmd/zona
ln -sf bin/zona ./zona
