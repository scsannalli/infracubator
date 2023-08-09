#Build go-application
FROM golang:latest as build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY pkg/http-server/server.go ./
RUN go build -o /go-docker-demo

#Run application
FROM build as run
EXPOSE 3333
CMD [ "/go-docker-demo" ]