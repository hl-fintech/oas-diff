### Go get dependecies and build ###
FROM golang:1.22.2 as builder
ENV PLATFORM docker
WORKDIR /go/src/app
COPY go.mod go.sum ./

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .

RUN VERSION=$(git describe --always --tags) && \
    CGO_ENABLED=0 GOOS=linux go build \
    -mod=readonly \
    -ldflags "-s -w -X github.com/hl-fintech/oas-diff/build.Version=${VERSION}"

### Create image ###
FROM alpine:3
WORKDIR /usr/bin
ENV PLATFORM github-action
COPY --from=builder /go/src/app/oasdiff .
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
