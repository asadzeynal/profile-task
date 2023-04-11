# Build stage
FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o grpc cmd/grpc/main.go
RUN go build -o gateway cmd/http/main.go
RUN go build -o swagger cmd/swagger/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/grpc .
COPY --from=builder /app/swagger .
COPY --from=builder /app/gateway .
COPY swaggerui/profile/v1/profile.swagger.json ./swaggerui/profile/v1/profile.swagger.json
COPY start.sh .

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082

CMD ["/app/start.sh"]
