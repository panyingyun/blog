#!/bin/bash
echo "start svn update"
svn update
echo "start svn update ok"

sleep 1

echo "start compiler"

cd michaelapp 

hugo 

cd ..

rm -rf public/
 
cp -r michaelapp/public/ ./

go generate

go build

sleep 1

tar czvf blog_amd64_v1.0.0-`date +%Y%m%d%H%M%S`.tar.gz  blog   appserver_prod.conf run.sh

echo "compiler & package ok"

echo "start update"

killall -9 blog

sleep 1

nohup ./blog -c appserver_prod.conf >homegoweb.log 2>&1 &

echo "kill & cp & run ok"
