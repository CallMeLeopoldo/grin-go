#!/usr/bin/env bash
set -euo pipefail

value="A,1,2,3,4,5,6,7,8,9,10,J,Q,K"
suit="clubs, diamonds, spades, cups"
json="["

for i in ${value//,/ }
do
    for j in ${suit//,/ }
    do
      json+=$(jq -n --arg value "$i" --arg suit "$j"  '{"suit":$suit, "value":$value}')
      json+=","
    done
done
json=$(echo ${json::-1})
json+="]"
echo $json > test.json