#!/bin/bash

kill -9 $(lsof -i:8080 -t)
kill -9 $(lsof -i:8081 -t)
kill -9 $(lsof -i:8082 -t)

kill -9 $(lsof -i:3000 -t)
kill -9 $(lsof -i:3001 -t)
kill -9 $(lsof -i:3002 -t)
