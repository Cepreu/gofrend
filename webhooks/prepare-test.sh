#!/bin/bash
curl "https://us-central1-f9-dialogflow-converter.cloudfunctions.net/handle-preparer-webhook" \
    --data-binary "@test_files/comparison_test.five9ivr"