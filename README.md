# Quick Farm

Problem Statement: In rural area farmers require different set of conventional and unconventional tools in one agriculture cycle and it's not possible for farmers to own each and every required  tool or rent it from lenders when required. Even if somebody own a tool it might become Non performing asset for them.

Objective: To allow farmers to lend and rent their famring tools using web platform.

# APIs

## User

- List Users : GET /user/{user_id}
- Get User Details : GET /user/{user_id}
- Edit User Profile : PUT /user/edit-profile/{user_id}
- User Login : POST /user/login
- User Registration : POST /user/register

## Equipment

- List Equipment : GET /equipments
- Get Equipment Details : GET /equipments/{equipment_id}
- Create Equipment : POST /equipments
- Update Equipment : PUT /user/{user_id}/equipments/{equipment_id}
- Delete Equipment : DELETE /user/{user_id}/equipments/{equipment_id}
- Get Equipment Owner : GET /owner/equipment/{equipment_id}
- Get User's Lended Equipment : GET /users/{user_id}/equipments/lended
- Rent Equipment : POST /users/{user_id}/equipments/{equip_id}/rent

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
