#!bin/bash

curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r --arg her $HERO_ID' .[] | select ( ( .id|tostring) == $hed) | (.connections.relatives | gsub("\n";"\\n"))'
