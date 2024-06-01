FROM golang:alpine AS build

WORKDIR /sso
COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o ./build/app ./cmd/sso

FROM alpine:latest

WORKDIR /app
COPY --from=build sso/build/app .
COPY --from=build sso/config .

ENTRYPOINT ["./app"]
CMD ["--config=config_local.yml"]
