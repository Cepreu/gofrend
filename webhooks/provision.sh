#!/bin/bash
gcloud endpoints services deploy openapi-functions.yaml
gcloud services enable servicemanagement.googleapis.com	servicecontrol.googleapis.com endpoints.googleapis.com