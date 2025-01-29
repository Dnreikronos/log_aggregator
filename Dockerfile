FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -o main ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 9090

CMD ["./main"]
