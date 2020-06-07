# Build binary
FROM golang:1.14-alpine as builder

ENV CGO_ENABLED 0

WORKDIR $GOPATH/src/app
COPY . .

# Unit tests
RUN go test ./...

# Build the binary.
RUN GOOS=linux \
    GOARCH=amd64 \
    go build \
        -ldflags="-w -s" \
        -o /go/bin/app \
        .

# Final Image
FROM alpine:latest

COPY --from=builder /go/bin/app /usr/local/bin/docker-semver
RUN chmod +x /usr/local/bin/docker-semver

ENTRYPOINT ["docker-semver"]
