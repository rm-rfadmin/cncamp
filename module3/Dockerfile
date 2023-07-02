FROM golang as builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM ubuntu:20.04

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x main

CMD ["./main"]
