package model

//sent business logic
type (
	User struct {
		ID           string `json:"id"`
		Username     string `json:"username"`
		Lastname     string `json:"lastname"`
		Password     string `json:"password"`
		Email        string `json:"email"`
		Role         string `json:"role"`
		ProfileImage string `json:"profile_image"`
	}
)
