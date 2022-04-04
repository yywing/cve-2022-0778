#! /bin/bash

tar -czh . | docker build -t server:latest -
docker run -it --rm -p 12345:12345 server:latest python server.py
