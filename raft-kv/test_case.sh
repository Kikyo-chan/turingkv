#!/bin/sh

ab -n 500 -c 10 -p post_data.txt -T 'application/json' http://10.0.2.5:8080/keys/aidwew/
