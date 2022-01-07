FROM golang:1.17.6-alpine

WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
COPY cmd cmd
COPY eveapi eveapi
COPY config.yml config.yml
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV CONFIG=./config.yml
RUN go test -c -o /go/bin/eveapi.test ./eveapi && eveapi.test -test.v
RUN go build -o /go/bin/eve-client cmd/cli/main.go
CMD [ "eve-client" ]
