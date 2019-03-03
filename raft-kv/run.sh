#!/bin/bash

# 运行leader进程
./output/bin/raft --bind 127.0.0.1:3000 --apiport :8080  --bootstrap true --datadir /tmp/data &

# 运行follower1进程
./output/bin/raft --bind 127.0.0.1:3001 --apiport :8081 --join 127.0.0.1:8080  --bootstrap false --datadir /tmp/data1 &

# 运行follower2进程
./output/bin/raft --bind 127.0.0.1:3002 --apiport :8082 --join 127.0.0.1:8080  --bootstrap false --datadir /tmp/data2 &

