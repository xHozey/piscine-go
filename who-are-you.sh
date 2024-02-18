#!/bin/bash

data=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json)

var_Name=$(echo "${data}" | jq '.[] | select(.id == 70) .name')

echo "${var_Name}"

