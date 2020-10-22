#!/bin/bash

mkdir tmp
rm -rf statik
LATEST=$(curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest |  grep tarball_url | cut -d : -f 2,3 |  tr -d \",)
curl -o swagger-ui.tar.gz -L $LATEST
tar xzvf swagger-ui.tar.gz -C tmp --strip-components 1
sed -i 's/"https:\/\/petstore.swagger.io\/v2\/swagger.json"/window.top.location.origin+"\/swagger.json"/g' tmp/dist/index.html
statik -src tmp/dist
rm -rf tmp
rm swagger-ui.tar.gz
