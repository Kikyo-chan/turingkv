#!/bin/bash


rm -rf /tmp/group_0_1
rm -rf /tmp/group_0_2
rm -rf /tmp/group_0_3

# run group pd

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:2000 --api_port 7080  --bootstrap true --data_dir /tmp/group_0_1 --group_id -1 &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:2001 --api_port 7081 --join 127.0.0.1:7080  --bootstrap false --data_dir /tmp/group_0_2 --group_id -1 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:2002 --api_port 7082 --join 127.0.0.1:7080  --bootstrap false --data_dir /tmp/group_0_3 --group_id -1 &


rm -rf /tmp/group0_1
rm -rf /tmp/group0_2
rm -rf /tmp/group0_3

# run group0

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:3000 --api_port 8080  --bootstrap true --data_dir /tmp/group0_1 --group_id 0 &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:3001 --api_port 8081 --join 127.0.0.1:8080  --bootstrap false --data_dir /tmp/group0_2 --group_id 0 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:3002 --api_port 8082 --join 127.0.0.1:8080  --bootstrap false --data_dir /tmp/group0_3 --group_id 0 &


rm -rf /tmp/group1_1
rm -rf /tmp/group1_2
rm -rf /tmp/group1_3

# run group1

# run leader process
nohup ./output/bin/raft --bind 127.0.0.1:4000 --api_port 9080  --bootstrap true --data_dir /tmp/group1_1 --group_id 1 &

# run follower1 process
nohup ./output/bin/raft --bind 127.0.0.1:4001 --api_port 9081 --join 127.0.0.1:9080  --bootstrap false --data_dir /tmp/group1_2 --group_id 1 &

# run follower2 process
nohup ./output/bin/raft --bind 127.0.0.1:4002 --api_port 9082 --join 127.0.0.1:9080  --bootstrap false --data_dir /tmp/group1_3 --group_id 1 &
