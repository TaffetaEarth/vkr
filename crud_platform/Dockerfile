FROM golang:alpine AS build

WORKDIR /app
COPY ./app .

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o app

FROM alpine:latest

WORKDIR /app
COPY --from=build app .

CMD ["./app"]
