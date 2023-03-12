# syntax=docker/dockerfile:1

FROM golang:1.20-alpine
RUN apk --no-cache add gcc g++ make git
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./

WORKDIR /app/cmd/api
RUN go build -o ../../bin

CMD [ "../../bin/api" ]

#run: docker run --network="host" api