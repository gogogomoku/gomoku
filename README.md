# Gomoku

⚪️⚫️⚪️⚫️

## Get it

```console
GO111MODULE=off go get -u github.com/gogogomoku/gomoku
```

## Development environment

### Install Golang dependencies

```console
export GO111MODULE=on
go get -v -d -u ./...
```

### Run server

```console
go run cmd/gomoku/main.go -s
```

### Run UI

```console
cd ./ui

npm install

npm run serve
```

### Performance

First, build it: `go build ./cmd/gomoku`

Run the binary and print out filename (slow):

```console
pprof_file="$(./gomoku 2> >(grep -o '[^[:space:]]*pprof') | tail -1)" && echo "CPU profile created: $pprof_file"

graph_pdf="./profgraph$(date '+%F_%T').pdf" && go tool pprof --pdf $pprof_file > $graph_pdf && open $graph_pdf
```

OR

Run from source:

```console
go run cmd/gomoku/main.go [-flags]

go tool pprof --pdf _FILE_ > file.pdf

open file.pdf
```

## Docker

Get a build up. You will need `docker` and `docker-compose`.

## Reminder

```console
docker-compose build --no-cache
docker-compose up --force-recreate
```
