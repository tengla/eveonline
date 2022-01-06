FROM golang:1.17.6-alpine

WORKDIR /usr/src/app
COPY go.mod .
COPY cmd cmd
COPY eveapi eveapi
RUN go build -o dist/client cmd/cli/main.go
CMD [ "./dist/client" ]
