FROM alpine:3.19

# Environment variables for Go
ENV GOLANG_VERSION=1.22.1 \
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
 && curl -LO https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && rm go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && mkdir -p /go/src /go/bin /go/pkg

# Install the 'op' tool
RUN go install lesiw.io/op@latest

# Create non-root user
RUN adduser -D -g '' runner
USER runner
WORKDIR /go

# Default command
CMD ["go", "version"]
