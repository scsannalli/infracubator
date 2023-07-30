#! /bin/sh

env GOOS=linux go build -o out/ ./...

docker build -t go-lang-service .

docker run -it --rm -p 3333:3333 go-lang-service