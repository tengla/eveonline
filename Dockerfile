FROM golang:1.17.6-alpine

WORKDIR /usr/src/app
COPY go.mod .
COPY cmd cmd
COPY eveapi eveapi
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go test -v ./eveapi && go build -o dist/client cmd/cli/main.go
CMD [ "./dist/client" ]
