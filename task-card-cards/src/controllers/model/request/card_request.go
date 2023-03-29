package request

type CardRequest struct {
	Title      string `json:"title" biding:"required"`
	Summary    string `json:"summary" biding:"required,max=2500"`
	DueDate    string `json:"due_date"`
	CardStatus string `json:"card_status" biding:"required"`
	UserId     int    `json:"user_id" biding:"required"`
}
