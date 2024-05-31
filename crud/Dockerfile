FROM golang:alpine AS build

WORKDIR /crud
COPY ./app ./app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o ./build/app ./app

FROM alpine:latest

WORKDIR /app
COPY --from=build crud/build/app .

CMD ["./app"]
