package user

type User struct {
	Id       string `json:"id"`
	Version  int64  `json:"version"`
	
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id    string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Version  int64  `json:"version" example:"1"`
	Email string `json:"email" example:"john.doe@example.com"`
}

type RegisterRequest struct {
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" example:"securepassword123"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
