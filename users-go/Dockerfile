FROM golang:1.15.6-alpine3.12 AS base
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

FROM base AS build
RUN go build -o /out/users cmd/main.go

FROM base AS unit-test
RUN go test -v -cover -bench=. ./...

FROM scratch AS bin
ENV USERS_GRPC_PORT=8080
COPY --from=build /out/users /
EXPOSE ${USERS_GRPC_PORT}
ENTRYPOINT ["/users"]