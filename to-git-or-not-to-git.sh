curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[] | select(.id == 170)' |jq -r '.name, .powerstats.power, .appearance.gender'





