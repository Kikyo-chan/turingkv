#!/bin/bash

rm -rf /tmp/data
rm -rf /tmp/data1
rm -rf /tmp/data2

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:3000 --apiport :8080  --bootstrap true --datadir /tmp/data &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:3001 --apiport :8081 --join 127.0.0.1:8080  --bootstrap false --datadir /tmp/data1 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:3002 --apiport :8082 --join 127.0.0.1:8080  --bootstrap false --datadir /tmp/data2 &
