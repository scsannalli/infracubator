
FROM golang:latest
WORKDIR /app
COPY go.mod /app/
RUN go mod download
COPY ./ /app/
RUN cd /app
RUN make build

#run app
EXPOSE 8080
CMD [ "/app/out/go-app" ]