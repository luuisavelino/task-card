package response

type CardResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	DueDate    string `json:"due_date"`
	CardStatus string `json:"card_status"`
	UserId     int    `json:"user_id"`
}
