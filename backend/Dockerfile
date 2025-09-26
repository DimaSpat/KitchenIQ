# syntax=docker/dockerfile:1.5

FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/go-build-cache \
    CGO_ENABLED=0 GOOS=linux go build \
    -o main \
    -v \
    -ldflags="-w -s" \
    ./src/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
