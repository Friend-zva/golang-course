FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN cd collector && go mod download
RUN cd collector && go build -o /bin/collector ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/collector /app/collector
COPY --from=builder /app/collector/config/local.yaml /app/config/local.yaml

RUN apk --no-cache add ca-certificates

EXPOSE 8081

CMD ["/app/collector"]
