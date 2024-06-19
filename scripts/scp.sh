#!/bin/bash

USR="ec2-user"
ROOT="$(cd "$(dirname $0)/.." && pwd)"

KEY="${ROOT}/infra/ec2.pem"
FILES="${ROOT}/build ${ROOT}/scripts"

scp -ri ${KEY} ${FILES} ${USR}@bot:~/
scp -ri ${KEY} ${FILES} ${USR}@attacker:~/
scp -ri ${KEY} ${FILES} ${USR}@middlebox:~/
