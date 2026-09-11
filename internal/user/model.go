package user

type User struct {
	Id       string `json:"id"`
	Version  int64  `json:"version"`
	
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Version  int64  `json:"version"`
	Email string `json:"email"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
