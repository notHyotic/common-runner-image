# Base: Alpine
FROM alpine:3.19

# Set environment variables for Go
ENV GOLANG_VERSION=1.22.1 \
    GOROOT=/usr/local/go \
    GOPATH=/go \
    PATH=/usr/local/go/bin:/go/bin:$PATH

# Install dependencies and Go
RUN apk add --no-cache curl git bash gcc musl-dev \
 && curl -LO https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && rm go${GOLANG_VERSION}.linux-amd64.tar.gz \
 && mkdir -p /go/src /go/bin /go/pkg

# Create a non-root user (optional but recommended)
RUN adduser -D -g '' runner
USER runner
WORKDIR /go

# Entrypoint (noop since GitHub Actions controls it)
CMD ["go", "version"]
