###############
# stage 1
###############
FROM golang:1.17.6-alpine3.15 AS BUILD_IMG

WORKDIR /usr/src/app
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o tmai-server .

###############
# stage 2
###############
FROM alpine:3.15

RUN mkdir /home/go/app && chown -R go:go /home/go/app

WORKDIR /home/go/app

COPY --from=BUILD_IMG /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=go:go --from=BUILD_IMG /usr/src/app/tmai-server .

USER go

EXPOSE 3000

ENTRYPOINT [/home/go/app/tmai-server]
