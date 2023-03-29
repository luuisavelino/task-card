package request

type UserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=40"`
	Userpass string `json:"userpass" binding:"required,min=6,containsany=!@#$%*"`
	Email    string `json:"email" binding:"required,email"`
	RoleId   int    `json:"role_id" binding:"required,min=1,max=2"`
}
