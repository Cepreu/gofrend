#!/bin/bash
rm go.mod
rm go.sum
export GO111MODULE=on
go mod init webhook
go mod tidy
gcloud components update
