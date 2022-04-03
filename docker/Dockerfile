FROM golang:1.18.0-bullseye AS base

WORKDIR /usr/src
COPY go.mod go.sum ./
RUN go mod download -x

###   dev   ###
FROM base as dev
WORKDIR /usr/src
RUN go install github.com/cespare/reflex@latest
COPY --from=base /go/pkg/mod /go/pkg/mod

###   release   ###
FROM base as compiler
WORKDIR /usr/src
COPY --from=base /go/pkg/mod /go/pkg/mod
COPY . ./
RUN go build -o ./app ./src

FROM debian:bullseye-slim as release
ENV GIN_MODE=release
WORKDIR /usr/src
EXPOSE 8080
COPY --from=compiler /usr/src/app ./
CMD ["./app"]
