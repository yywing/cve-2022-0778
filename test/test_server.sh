#! /bin/bash

docker run -it --rm python:3.7.1  python -c "import http.client;http.client.HTTPSConnection('host.docker.internal', 12345).request('GET', '/')"
