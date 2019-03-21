#!/bin/sh

echo -e "SET KEY (some-key) TEST \n"

curl 'http://127.0.0.1:8000/keys/turing-key-dcdwcbkwqcbhekwqbckejwqhchb/' -H 'Content-Type: application/json' -d '{"value": "hello turingkv1"}'

echo -e "\nGET KEY TEST \n"

echo -e "GET VALUE OF some-key \n"

curl 'http://127.0.0.1:8080/keys/some-key/'

echo -e "AB TEST \n"

ab -n 100 -c 100 http://127.0.0.1:8080/keys/some-key/