# Marketplace Go Backend

A microservices-based marketplace backend built with Go, following a similar architecture to the gojobber-backend project.

## Project Structure

```
marketpace-go-backend/
├── services/
│   ├── 1-gateway/          # API Gateway service
│   ├── 2-notification/     # Notification service
│   ├── 3-auth/             # Authentication service
│   ├── 4-user/             # User management service
│   ├── 5-product/          # Product management service
│   ├── 6-order/            # Order management service
│   ├── 7-payment/           # Payment processing service
│   ├── 8-review/            # Review service
│   └── common/              # Shared code and generated protobuf
├── protobuf/               # Protocol buffer definitions
├── go.mod                  # Go module file
├── Makefile                # Build and run commands
└── .gitignore             # Git ignore rules
```

## Services

1. **Gateway Service**: Main API gateway that routes requests to appropriate microservices
2. **Notification Service**: Handles email notifications and other messaging
3. **Auth Service**: User authentication and authorization
4. **User Service**: User profile management (buyers and sellers)
5. **Product Service**: Product catalog management
6. **Order Service**: Order processing and management
7. **Payment Service**: Payment processing integration
8. **Review Service**: Product and seller reviews

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Generate protobuf code:
```bash
make proto-auth
make proto-user
make proto-notification
make proto-product
make proto-order
make proto-payment
make proto-review
```

3. Set up environment variables in `.env` file

4. Run services:
```bash
make run-gateway
make run-auth
make run-user
# ... etc
```

## Development

This project is being developed over 10 days. Today (Day 1) focuses on setting up the basic project structure.

## License

[Add your license here]
