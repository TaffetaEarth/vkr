FROM golang:alpine AS build

WORKDIR /streamer
COPY ./app ./app

COPY go.mod .
COPY go.sum .

RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod vendor
RUN go build -o ./build/app ./app

FROM alpine:latest

WORKDIR /app
COPY --from=build streamer/build/app .

CMD ["./app"]


