FROM golang:alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/hello

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/hello .

CMD ["./hello"]
