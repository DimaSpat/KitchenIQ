# Stage 1: The Build Stage
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main -v -ldflags="-w -s" ./src/main.go


FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 5000

ENTRYPOINT ["./main"]
