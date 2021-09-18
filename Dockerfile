FROM golang:1.17.1-alpine

RUN apk update && apk --no-cache add git build-base

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["fresh"]
