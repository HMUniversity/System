package member_mod

type JoinRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

type JoinResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
