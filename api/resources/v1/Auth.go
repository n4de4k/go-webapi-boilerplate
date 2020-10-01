package v1

type (
	SignInRequest struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	SignInResponse struct {
		Token string `json:"token"`
	}
)
