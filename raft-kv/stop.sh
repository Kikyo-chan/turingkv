#!/bin/bash

rm -rf nohup.out

kill -9 $(lsof -i:8080 -t)
kill -9 $(lsof -i:8081 -t)
kill -9 $(lsof -i:8082 -t)

kill -9 $(lsof -i:3000 -t)
kill -9 $(lsof -i:3001 -t)
kill -9 $(lsof -i:3002 -t)

kill -9 $(lsof -i:9080 -t)
kill -9 $(lsof -i:9081 -t)
kill -9 $(lsof -i:9082 -t)

kill -9 $(lsof -i:4000 -t)
kill -9 $(lsof -i:4001 -t)
kill -9 $(lsof -i:4002 -t)

kill -9 $(lsof -i:7080 -t)
kill -9 $(lsof -i:7081 -t)
kill -9 $(lsof -i:7082 -t)

kill -9 $(lsof -i:2000 -t)
kill -9 $(lsof -i:2001 -t)
kill -9 $(lsof -i:2002 -t)