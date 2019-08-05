#!/bin/bash
curl \
    -d "SCRIPT_NAME=sk_text&CAMPAIGN_NAME=sergei_inbound&USERNAME=sam2@007.com&TEMPORARY_PASSWORD=pwd1234567" \
    -H "ACCESS_TOKEN: GMWJGSAPGATLMODYLUMGUQDIWWMNPPQI" \
    "https://us-central1-f9-dialogflow-converter.cloudfunctions.net/handle-preparer-webhook"
    