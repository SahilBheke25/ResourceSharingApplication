# ResouceSharingApplication
This repo contains my Resource Sharing Application project.
Objective: To allow farmers to lend and rent their famring tools using web platform.

# APIs

## User

List Users : GET http://localhost:8080/user/{user_id}
Get User Details : GET http://localhost:8080/user/{user_id}
Edit User Profile : PUT http://localhost:8080/user/edit-profile/{user_id}
User Login : POST http://localhost:8080/user/login
User Registration : POST http://localhost:8080/user/register

## Equipment

List Equipment : GET http://localhost:8080/equipments
Get Equipment Details : GET http://localhost:8080/equipments/{equipment_id}
Create Equipment : POST http://localhost:8080/equipments
Update Equipment : PUT http://localhost:8080/user/{user_id}/equipments/{equipment_id}
Delete Equipment : DELETE http://localhost:8080/user/{user_id}/equipments/{equipment_id}
Get Equipment Owner : GET http://localhost:8080/owner/equipment/{equipment_id}
Get User's Lended Equipment : GET http://localhost:8080/users/{user_id}/equipments/lended
Rent Equipment : POST http://localhost:8080/users/{user_id}/equipments/{equip_id}/rent

# Project File Structure
```
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
```
