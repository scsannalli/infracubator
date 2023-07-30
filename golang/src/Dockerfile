FROM golang:latest
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY pkg/http-server/server.go ./

RUN go build -o /go-docker-demo

EXPOSE 3333

CMD [ "/golang/src" ]