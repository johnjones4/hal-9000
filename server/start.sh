#!/bin/sh

curl $SERVICE_CONFIG_URL > $SERVICE_CONFIG
curl $INTENT_DESCRIPTIONS_URL > $INTENT_DESCRIPTIONS

/app/hal9000