package response

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Email    string `json:"email" validate:"nonzero, regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"`
	RoleId   int    `json:"role_id" validate:"roles"`
}
