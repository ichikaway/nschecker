#!/bin/sh

GOFILE='tld_servers.go'

echo 'package rootmaster

// DNS gTLD servers Array.
func getTldServers() map[string]string {
	servers := make(map[string]string, 3)
' > $GOFILE

cat tld-servers.txt | awk -F ". " '{print "\tservers[\"" $1 "\"] = \"" $2 "\""}' >> $GOFILE

echo '
	return servers
}' >> $GOFILE
