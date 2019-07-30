# Gomoku 

⚪️⚫️⚪️⚫️
## Development environment

### Install Golang dependencies

```
$ export GO111MODULE=on
$ go gen -v -d -u ./...
```

### Docker

Run `docker build .` from repository root.

Run `docker run -p 4242:4242 --name gomoku` from repository root.