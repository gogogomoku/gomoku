# Gomoku 

⚪️⚫️⚪️⚫️

## Get it

```
GO111MODULE=off go get -u github.com/gogogomoku/gomoku
```

## Development environment

### Install Golang dependencies

```
$ export GO111MODULE=on
$ go get -v -d -u ./...
```

### Run server

```
$ go run cmd/gomoku/main.go -s
```

### Run UI

```
$ cd ./ui
$ npm install
$ npm run serve
```

## Docker
Get a build up. You will need `docker` and `docker-compose`.

#### Reminder
```
docker-compose build --no-cache
docker-compose up --force-recreate
```
