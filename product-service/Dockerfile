############################
# STEP 1 build executable binary
############################
FROM golang:1.17-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o product-service

############################
# STEP 2 build a small image
############################
FROM alpine
RUN apk --no-cache add curl
WORKDIR /app
COPY --from=builder /app/product-service /app
CMD ["./product-service"]