#!/bin/bash

# JSON_FNAME=$1
JSON_FNAME=result.json
JSON_PATH='plaintext'
HEX_PER_LINE=16

i=1
for d in $(cat ${JSON_FNAME} | jq -r ".${JSON_PATH}[]"); do
  printf "0x%02x, " $d
  if [[ $(( i % $HEX_PER_LINE )) == 0 ]]; then
    echo
  fi
  i=$(( ${i} + 1 ))
done
