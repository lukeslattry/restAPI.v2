#!/bin/bash

# Run this curl command to test if we can post some data
# Go to: http://localhost/todos to see the response from the API

curl -H "Content-Type: applictaion/json" -d '{ "name": "Now Todo" }' http://localhost:8080/todos