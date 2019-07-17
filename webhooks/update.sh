#!/bin/bash
rm go.mod
rm go.sum
export GO111MODULE=on
go mod init
go mod tidy
gcloud components update
gcloud functions deploy handle-fulfiller-webhook --entry-point=HandleFulfillerWebhook --runtime=go111 --trigger-http
gcloud functions deploy handle-preparer-webhook --entry-point=HandlePreparerWebhook --runtime=go111 --trigger-http