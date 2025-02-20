package models

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Pincode    int    `json:"pincode"`
	Uid        int    `json:"uid"`
}

type UserProfile struct {
	Username   string `json:"username"`
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
}

type UserCredentials struct {
	Password string `json:"password"`
}
