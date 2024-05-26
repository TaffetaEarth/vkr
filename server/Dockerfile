FROM golang:alpine AS build

WORKDIR /app
COPY ./app .

COPY go.mod .
COPY go.sum .
COPY classic.mp3 .

RUN go mod download
RUN go build -o /app/hello

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/hello .

CMD ["./hello"]
