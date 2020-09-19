#!/bin/sh

curl https://www.internic.net/domain/root.zone > root.zone 
grep "\tNS\t" root.zone | awk '!/^\./ {print $1, $5}' > ./tld-servers.txt
