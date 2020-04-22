#!/bin/sh

env GOOS=linux GOARCH=amd64 go build -o ./pritunl-pass-hash main.go