package dto

type (
	UserResponse struct {
		Email string `json:"email"`
		FullName string `json:"fullName"`
		Password string `json:"password"`
	}

	GetUserRequest struct {
		Email string `json:"email"`
	}
)
