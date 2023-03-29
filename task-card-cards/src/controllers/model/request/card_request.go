package request

type CardRequest struct {
	Title      string `json:"title" binding:"required"`
	Summary    string `json:"summary" binding:"required,max=2500"`
	DueDate    string `json:"due_date"`
	CardStatus string `json:"card_status" binding:"required"`
	UserId     int    `json:"user_id" binding:"required"`
}

type ActionRequest struct {
	UserId int `json:"user_id" binding:"required"`
}
