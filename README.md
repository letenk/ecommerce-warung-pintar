# Microservice

Microservice to test Warung Pintar

## Documentation

[Swagger Documentation](https://app.swaggerhub.com/apis-docs/DARMAWANRIZKY43/warung_pintar/1.0.0)

[Postman Documentation](https://documenter.getpostman.com/view/12132212/UzBqnQDz)

[ERD Documentation](https://dbdiagram.io/d/62b10c8f69be0b672c085fcb)

## Installation

Clone the project

```bash
  git clone https://github.com/jabutech/ecommerce-warung-pintar.git
```

Go to the project directory

```bash
  cd ecommerce-warung-pintar
```

Start the server

```bash
  docker-compose -f docker-compose.prod.dev up -d
```

## Run Test Service Auth

1. Go to folder `auth-service`.

2. Export environtment variable into terminal

```bash
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=secret
export DB_PASSWORD=secret
export DB_NAME=users
export DB_NAME_TEST=user_tests
export SECRET_JWT=w4run6P1Nt42
```

3. Run command

```bash
go test -v ./...
```

## Run Test Service Product

1. Go to folder `product-service`.

2. Export environtment variable into terminal

```bash
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=secret
export DB_PASSWORD=secret
export DB_NAME=products
export DB_NAME_TEST=product_tests
```

3. Run command

```bash
go test -v ./...
```

## Additional Goals

For the next stage, will be built:

- [x] Carts service for to add the product selected by the user
- [x] Orders service to process orders to create invoices

But I'm running out of time, if you have free time can you give me some feedback on what to input. So that it can be a reference for me in the future to develop.

Thank you and best regards.
