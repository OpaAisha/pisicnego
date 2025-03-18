#!/bin/bash

curl -s "https://01.tomorrow-school.ai/assets/superhero/all.json" | jq -r  ".[] | select(.id == $HERO_ID) | .connections.relatives" | sed ':a;N;$!ba;s/\n/\\n/g'
