############################
# STEP 1 build executable binary
############################
FROM golang:1.17-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o api-gateway

############################
# STEP 2 build a small image
############################
FROM alpine
WORKDIR /app
COPY --from=builder /app/api-gateway /app
CMD ["./api-gateway"]