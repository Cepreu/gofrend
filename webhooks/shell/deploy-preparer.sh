#!/bin/bash
gcloud functions deploy handle-preparer-webhook --entry-point=HandlePreparerWebhook --runtime=go111 --trigger-http