#!/bin/bash

mkdir logs
./output/bin/raft-proxy --api_port :9988 --group_count 2
