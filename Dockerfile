FROM golang:1.17.6-alpine

WORKDIR /usr/src/app
COPY go.mod .
COPY cmd cmd
COPY eveapi eveapi
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go test -c -o /go/bin/eveapi.test ./eveapi && eveapi.test -test.v
RUN go build -o /go/bin/eve-client cmd/cli/main.go
CMD [ "eve-client" ]
