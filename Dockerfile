# Multi stage buils FTW!: https://blog.docker.com/2017/07/multi-stage-builds/
FROM golang:1.20 AS builder
MAINTAINER Leandro López (inkel) <leandro@citrusbyte.com>

WORKDIR /go/src/lruc

ADD ["go.mod", "main.go", "."]

# Compile an optimized version to be deployed as a Docker container
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -a -ldflags='-w -s -extldflags "-static"' .

# Final container image
FROM scratch
COPY --from=builder /go/src/lruc/lruc /lruc

EXPOSE 8080

ENTRYPOINT ["/lruc"]
