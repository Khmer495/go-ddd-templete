FROM golang:1.17.5-bullseye

WORKDIR /go_templete/app

COPY ./script/install_tools.sh ./script/install_tools.sh

RUN ./script/install_tools.sh

COPY go.mod .

RUN go mod tidy
