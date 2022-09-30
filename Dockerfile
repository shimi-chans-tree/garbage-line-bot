FROM golang:1.17.7-alpine

RUN apk update &&  apk add git

RUN go get github.com/cosmtrek/air@v1.40.4

WORKDIR /app

CMD ["air", "-c", ".air.toml"]

