#Build go-application
FROM golang:latest
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY pkg/http-server/server.go ./
RUN go build -o /go-docker-demo

#Run application
EXPOSE 3333
CMD [ "/go-docker-demo" ]