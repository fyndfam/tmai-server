## tmai-server

![Build and Test](https://github.com/fyndfam/tmai-server/actions/workflows/build_and_test.yml/badge.svg)
[![deploy](https://github.com/fyndfam/tmai-server/actions/workflows/deploy.yml/badge.svg)](https://github.com/fyndfam/tmai-server/actions/workflows/deploy.yml)

tmai-server is the backend server for tmai written in Go

### How to get started

- make sure you have go install. At the time of writing this readme, we're using v1.17 of Go

#### To run the server locally

- you need to have docker, docker-compose and a container running mongodb, spin it up by the following command

```sh
docker-compose -f docker-compose.yml up -d mongo
```

- then install depedencies by

```sh
go get
```

- then you can run the server and it will listen at port 3000

```sh
MONGODB_URL=mongodb://tmai:password@localhost:27017/tmai go run main.go
```


#### To run the test

- same as above, you'll need to make sure that the mongo container is up
- run the test with this command

```sh
MONGODB_URL=mongodb://tmai:password@localhost:27017/tmai go test -v ./...
```
