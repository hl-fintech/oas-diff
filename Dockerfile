### Go get dependecies and build ###
FROM golang:1.22.2 as builder
ENV PLATFORM docker
WORKDIR /go/src/app
COPY go.mod go.sum ./

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -mod=readonly

### Create image ###
FROM alpine:3
USER root
WORKDIR /usr/bin
ENV PLATFORM github-action
COPY --from=builder /go/src/app/oasdiff .

COPY entrypoint.sh /entrypoint.sh
RUN echo "$GITHUB_OUTPUT"

ENTRYPOINT ["/entrypoint.sh"]
