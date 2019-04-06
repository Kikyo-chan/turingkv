#!/bin/bash

rm -rf logs
mkdir logs

rm -rf /nvmessd/group_0_1
rm -rf /nvmessd/group_0_2
rm -rf /nvmessd/group_0_3

# run group pd

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:2000 --api_port 7080 --rpc_port :7000  --bootstrap true --data_dir /nvmessd/group_0_1 --group_id -1 &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:2001 --api_port 7081 --rpc_port :7001 --join 127.0.0.1:7080  --bootstrap false --data_dir /nvmessd/group_0_2 --group_id -1 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:2002 --api_port 7082 --rpc_port :7002 --join 127.0.0.1:7080  --bootstrap false --data_dir /nvmessd/group_0_3 --group_id -1 &

rm -rf /nvmessd/group0_1
rm -rf /nvmessd/group0_2
rm -rf /nvmessd/group0_3

# run group 0

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:3000 --api_port 8080 --rpc_port :8000  --bootstrap true --data_dir /nvmessd/group0_1 --group_id 0 &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:3001 --api_port 8081 --rpc_port :8001 --join 127.0.0.1:8080  --bootstrap false --data_dir /nvmessd/group0_2 --group_id 0 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:3002 --api_port 8082 --rpc_port :8002 --join 127.0.0.1:8080  --bootstrap false --data_dir /nvmessd/group0_3 --group_id 0 &

rm -rf /nvmessd/group1_1
rm -rf /nvmessd/group1_2
rm -rf /nvmessd/group1_3

# run group 1

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:4000 --api_port 9080 --rpc_port :9000  --bootstrap true --data_dir /nvmessd/group1_1 --group_id 1 &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:4001 --api_port 9081 --rpc_port :9001 --join 127.0.0.1:9080  --bootstrap false --data_dir /nvmessd/group1_2 --group_id 1 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:4002 --api_port 9082 --rpc_port :9002 --join 127.0.0.1:9080  --bootstrap false --data_dir /nvmessd/group1_3 --group_id 1 &
