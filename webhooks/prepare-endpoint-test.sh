#!/bin/bash
curl \
    --header "content-type:application/x-www-form-urlencoded" \
    --data-urlencode "SCRIPT_NAME=sk_text&CAMPAIGN_NAME=sergei_inbound&USERNAME=sam2@007.com&TEMPORARY_PASSWORD=pwd1234567" \
    "https://f9-dialogflow-converter-endpoint-vi2zkckh3a-uc.a.run.app/handle-preparer-webhook"
    