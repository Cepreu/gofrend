#!/bin/bash
curl \
    -d "SCRIPT_NAME=ZPizza&CAMPAIGN_NAME=sergei_inbound&USERNAME=sam2@007.com&TEMPORARY_PASSWORD=pwd1234567" \
    "https://us-central1-f9-dialogflow-converter.cloudfunctions.net/handle-preparer-webhook"
    