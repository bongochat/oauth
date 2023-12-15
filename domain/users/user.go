package users

type User struct {
	Id          int64  `json:"id"`
	PhoneNumber string `json:"phone_number"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
