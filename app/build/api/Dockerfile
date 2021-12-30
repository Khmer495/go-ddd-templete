FROM golang:1.17.5-bullseye

WORKDIR /go/src/app

COPY . .
RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w' \
    ./cmd/api/v1/main.go

FROM alpine:3.14.0

WORKDIR /app

COPY some_file some_file
COPY --from=builder /go/bin/main main

ENTRYPOINT ["./main"]
