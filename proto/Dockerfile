FROM alpine:latest as base

RUN apk add --no-cache curl
RUN apk add --no-cache go
RUN apk add --no-cache protoc
RUN apk add --no-cache git

# Install the protocol compiler plugins for Go
RUN GO111MODULE=on go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25 \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Install buf.build
RUN VERSION=0.33.0 && \
    curl -sSL https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m) \
        -o /usr/local/bin/buf && \
    chmod +x /usr/local/bin/buf

WORKDIR proto
COPY . .

FROM base as lint
RUN export PATH=$PATH:$(go env GOPATH)/bin && buf check lint

FROM base as gen
RUN export PATH=$PATH:$(go env GOPATH)/bin && buf generate
RUN cd gen/go && go mod init github.com/demeero/shop-sandbox/proto/gen/go && go mod tidy
