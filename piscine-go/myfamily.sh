#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives' |tr -d '""'
