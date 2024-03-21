FROM golang:1.22.1-alpine as base

FROM base as dev

WORKDIR /opt/app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
