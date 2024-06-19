#!/bin/bash

echo -n "" > iodine.csv
for json in $(find ~/Downloads/dns-tunnel/ztmb/ -name '*.json'); do
  echo -n "${json}," >> iodine.csv
  ./build/ztmb-wo-zkp ${json} | grep '0x20 Modified/Total' | awk '{print $3}' | tr '/' ',' | tee -a iodine.csv
done
