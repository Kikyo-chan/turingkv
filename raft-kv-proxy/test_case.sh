#!/bin/bash

curl 'http://127.0.0.1:9988/keys/a-key/' -H 'Content-Type: application/json' -d '{"value": "hello turingkv1"}'

curl 'http://127.0.0.1:9988/keys/a-key/'
