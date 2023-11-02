package model

type User struct {
	Tel      string `json:"tel"`
	UUID     string
	Password string `json:"password"`
}

type UserResponse struct {
	UUID           string `json:"uuid"`
	HashedPassword string `json:"hashed_password"`
	IsPinSet       bool   `json:"is_pin_set"`
	IsProfileSet   bool   `json:"is_profile_set"`
}

type UserProfile struct {
	UUID      string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       string `json:"age"`
}

type Pin struct {
	Tel string `json:"tel"`
	Pin string `json:"pin"`
}

type SetNewPin struct {
	Tel    string `json:"tel"`
	Pin    string `json:"pin"`
	NewPin string `json:"new_pin"`
}

type UpdatePassword struct {
	Tel         string `json:"tel"`
	Otp         string `json:"otp"`
	Password    string `json:"current_password"`
	NewPassword string `json:"new_password"`
}
