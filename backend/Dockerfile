FROM golang

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 5000
ENV PORT=:5000

RUN go build ./src/main.go

CMD ["./main"]