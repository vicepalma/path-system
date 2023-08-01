#!/bin/bash
# echo "Export Go path"
# export PATH="$PATH:$(go env GOPATH)/bin"

echo "Generate code"
gen --sqltype=postgres --database path-system --json --gorm --mod --out ./system-orm --json-fmt=snake --generate-dao --generate-proj --overwrite --connstr "host=localhost port=5432 user=path_admin password=path_secret dbname=path-system sslmode=disable"

# gen --sqltype=postgres \
# --database path-system \
# --json \
# --gorm \
# --mod \
# --out ./example \
# --json-fmt=snake \
# --generate-dao \
# --generate-proj \
# --overwrite \
# --connstr "host=localhost port=5432 user=path_admin password=path_secret dbname=path-system sslmode=disable"