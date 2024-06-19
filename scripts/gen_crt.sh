#!/bin/bash

cat << EOF > proxy.conf
[req]
distinguished_name = cert_dn
x509_extensions = v3_req
prompt = no

[cert_dn]
C = ${COUNTRY:=KR}
ST = ${STATE:=Seoul}
L = ${LOCALITY:=Seoul}
O = ${ORGANIZATION:=Seoul National University}
OU = ${ORGANIZATION_UNIT:=Computer Science and Engineering}
CN = ${COMMON_NAME:=attacker.ztmb.io}

[v3_req]
keyUsage = keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.0 = ${DOMAIN:=attacker.ztmb.io}
IP.0 = ${IP_ADDR:=172.16.20.10}
EOF

openssl ecparam -genkey -name prime256v1 -out proxy.key
openssl req -x509 -key proxy.key -config proxy.conf -out proxy.crt -days 365
