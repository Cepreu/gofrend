#!/bin/bash
gcloud functions deploy handle-fulfiller-webhook --entry-point=HandleFulfillerWebhook --runtime=go111 --trigger-http