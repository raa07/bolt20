FROM golang:1.13-alpine

WORKDIR /app

RUN apk update
RUN apk add git


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

EXPOSE 80

CMD [ "/main" ]