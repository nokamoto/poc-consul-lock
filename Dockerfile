FROM golang:1.13.6-alpine3.11 AS build

RUN apk update && apk add git

WORKDIR /src

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go install .

FROM alpine:3.11

RUN apk update && apk add --no-cache ca-certificates

COPY --from=build /go/bin/poc-consul-lock /usr/local/bin/poc-consul-lock

ENTRYPOINT [ "poc-consul-lock" ]
