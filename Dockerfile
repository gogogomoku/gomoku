# Stage 1
FROM golang:1.12.6-alpine3.10 AS build
ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/gogogomoku/gomoku/
COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download
RUN go build -o /bin/gomoku ./cmd/gomoku/

# Stage 2
FROM alpine as run
COPY --from=build /bin/gomoku /bin/gomoku
ENTRYPOINT ["/bin/gomoku"]