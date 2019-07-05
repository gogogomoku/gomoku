# go-go-GOMOKU!

## Working on Gomoku

### Set up development workspace
```
mkdir $GOPATH/github.com/gogogomoku
cd $GOPATH/github.com/gogogomoku
git clone https://github.com/gogogomoku/gomoku.git
```

### Compile and run in development
```
go run cmd/gomoku/main.go -s
```

## Building

### Build in project root
```
go build cmd/gomoku/main.go # creates the executable ./main
```

### Building in $GOPATH/bin

(Faster, caches submoduels in $GOPATH/pkg)
```
go install -v ./... # creates the executable $GOPATH/bin/gomoku
```

## Misc

### Tidying up dependency requirements

For our purposes, this will remove unused deps from `go.mod`.
```
export GO111MODULE=on
go mod tidy
```

#### Vendoring (should we need it)

This should not be necessary unless we want to incorporate vendor libraries into our codebase. It will change our project's module directory to `vendor` and use its contents to source our dependencies.
```
go mod vendor
go build -mod vendor cmd/gomoku/main.go # build using code in <project_root>/vendor
```