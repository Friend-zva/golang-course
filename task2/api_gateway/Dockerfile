
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN cd api_gateway && go mod download
RUN cd api_gateway && go build -o /bin/api_gateway ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/api_gateway /app/api_gateway
COPY --from=builder /app/api_gateway/config/local.yaml /app/config/local.yaml

EXPOSE 8080

CMD ["/app/api_gateway"]
