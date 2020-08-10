FROM golang:1.14.6-alpine3.11
RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 8080
EXPOSE 8888
RUN go mod download
RUN go run cmd/lineProcessor/main.go
#RUN go build -v ./cmd/lineProcessor -o lineProcessor
