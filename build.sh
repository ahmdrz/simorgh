#!/bin/bash

set -e

TMP="$(mktemp -d -t simorgh.XXXX)"

function rmtemp {
	rm -rf "$TMP"
}
trap rmtemp EXIT

CURRENT="$(pwd)"
SIMORGH="$TMP/src/github.com/ahmdrz/simorgh"

export GOPATH="$TMP":$GOPATH
for file in `find . -type f`; do
	if [[ "$file" != "." && "$file" != ./.git* && "$file" != ./data* && "$file" != ./summitdb-* ]]; then
		mkdir -p "$SIMORGH/$(dirname "${file}")"
		cp -P "$file" "$SIMORGH/$(dirname "${file}")"
	fi
done

cd $SIMORGH
cd server && go build -i -o $CURRENT/simorgh-server
cd ../client && go build -i -o $CURRENT/simorgh-client