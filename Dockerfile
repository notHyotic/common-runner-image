FROM alpine:3.19

# Environment variables for Go
ENV GOLANG_VERSION=1.24.3 \
    GOROOT=/usr/local/go \
    GOPATH=/go \
    PATH=/usr/local/go/bin:/go/bin:$PATH

# Install system packages, Go, and Docker CLI
RUN apk add --no-cache \
    bash \
    curl \
    git \
    gcc \
    musl-dev \
    docker-cli \
    ca-certificates \
    nodejs \
    npm \
 && curl -LO https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && rm go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && mkdir -p /go/src /go/bin /go/pkg

# Install the 'op' tool
RUN go install lesiw.io/op@latest

# Install the 'golangci-lint' tool
RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6

# Default command
CMD ["go", "version"]
