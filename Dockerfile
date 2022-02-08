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

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN mkdir -pv /home/appuser/app && chown -R appuser:appgroup /home/appuser/app

WORKDIR /home/appuser/app

COPY --from=BUILD_IMG /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=appuser:appgroup --from=BUILD_IMG /usr/src/app/tmai-server .

USER appuser

EXPOSE 3000

ENTRYPOINT [/home/appuser/app/tmai-server]
