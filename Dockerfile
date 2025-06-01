FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch
COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]