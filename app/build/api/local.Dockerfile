FROM golang:1.17.5-bullseye

WORKDIR /go/src/app

COPY . .
RUN go mod download && \
    go install github.com/cosmtrek/air@v1.27.3

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

CMD ["air", "-c", ".air.toml"]
