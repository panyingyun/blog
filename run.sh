#!/bin/bash

echo "start update"

killall -9 blog

sleep 1

nohup ./blog -c appserver_prod.conf >homegoweb.log 2>&1 &

echo "kill & cp & run ok"
