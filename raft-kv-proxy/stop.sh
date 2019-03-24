#!/bin/bash

kill -9 $(lsof -i:9988 -t)
