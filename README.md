# Quick Farm

Problem Statement: In rural area farmers require different set of conventional and unconventional tools in one agriculture cycle and it's not possible for farmers to own each and every required  tool or rent it from lenders when required. Even if somebody own a tool it might become Non performing asset for them.

Objective: To allow farmers to lend and rent their famring tools using web platform.

# Project Documentation
<a href="https://docs.google.com/document/d/1-nHlaTeGwmPCY1gzvjb_Zw78IeqDkPvhQHNOcsjo-u0/edit?usp=sharing">Doc Link</a>

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
  

# Postman Collection 
<a href="https://solar-star-172287.postman.co/workspace/My-Workspace~db048750-3973-4142-8e4d-36ad4cd7cf0b/collection/41461760-c03c7feb-fb93-4d63-996d-767d1868c8c5?action=share&creator=41461760">Quick-Share</a>


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
