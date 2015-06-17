#!/bin/bash

# Run this curl command to test if we can post some data

curl -H "Content-Type: applictaion/json" -d '{ "name": "Now Todo" }' http://localhost:8080/todos

