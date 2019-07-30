FROM golang:1.12.6-alpine3.10 AS build
ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/gogogomoku/gomoku/

COPY . .

RUN go get -v -d -u ./...
RUN go build -o /bin/gomoku ./cmd/gomoku/
FROM alpine as run

COPY --from=build /bin/gomoku /bin/gomoku
ENTRYPOINT ["/bin/gomoku"]
CMD ["-s"]