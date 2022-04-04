FROM python:3.7.1

COPY server.py /app/server.py
ADD server.crt /app/server.crt
ADD server.key /app/server.key

WORKDIR /app
EXPOSE 12345