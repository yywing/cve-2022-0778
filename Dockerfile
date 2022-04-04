FROM golang:1.17-alpine as builder


COPY . /app
WORKDIR /app
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN CGO_ENABLED=0 go build -o /app/badserver ./server
RUN CGO_ENABLED=0 go build -o /app/badclient ./client


FROM alpine

COPY --from=builder /app/badserver /bin/badserver
COPY --from=builder /app/badclient /bin/badclient

ENTRYPOINT [ "badserver" ]
