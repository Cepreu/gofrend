#!/bin/bash
curl \
    -F SCRIPT_NAME=sk_text \
    -F CAMPAIGN_NAME="sergei_inbound" \
    -F USERNAME="sam2@007.com" \
    -F TEMPORARY_PASSWORD="pwd1234567" \
    "https://us-central1-f9-dialogflow-converter.cloudfunctions.net/handle-preparer-webhook"
    