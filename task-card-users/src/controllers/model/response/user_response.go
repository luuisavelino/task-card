package response

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
}
