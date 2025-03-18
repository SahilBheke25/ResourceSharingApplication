# ResouceSharingApplication
This repo contains my Resource Sharing Application project.
Objective: To allow farmers to lend and rent their famring tools using web platform.

# Project File Structure
</br>
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── dependencies.go
│   │   ├── equipment
│   │   │   ├── handler.go
│   │   │   └── service.go
│   │   ├── rental
│   │   │   ├── handler.go
│   │   │   └── service.go
│   │   ├── routes.go
│   │   ├── user
│   │   │   ├── handler.go
│   │   │   ├── mocks
│   │   │   │   └── Service.go
│   │   │   ├── service.go
│   │   │   └── user_test.go
│   │   └── utils
│   │       ├── hashing.go
│   │       └── response_writter.go
│   ├── config
│   │   └── config.go
│   ├── models
│   │   ├── billing.go
│   │   ├── equipment.go
│   │   ├── rental.go
│   │   └── user.go
│   ├── pkg
│   │   ├── apperrors
│   │   │   └── erros.go
│   │   └── middleware
│   │       ├── auth.go
│   │       └── request_auth.go
│   └── repository
│       ├── connection.go
│       ├── equipment.go
│       ├── rental.go
│       └── user.go
└── README.md
