FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

COPY . .

RUN go mod download 2>/dev/null || true

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

ENV PORT=8080
ENV NODE_API_URL=http://node-api:3000

CMD ["./main"]
