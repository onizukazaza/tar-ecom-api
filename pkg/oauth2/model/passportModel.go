package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID       string `json:"id"`
	Role         string `json:"role"`
	Token        string `json:"token"`
	Username     string `json:"username,omitempty"`     
	ProfileImage string `json:"profile_image,omitempty"` 
}
