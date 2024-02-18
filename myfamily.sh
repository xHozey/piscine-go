
#!/bin/bash

data=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json)

var_Name=$(echo "${data}" | jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber))' | jq '.connections.relatives')
background_Result=$(echo "${var_Name}" | sed 's/"//g')
echo "${background_Result}"