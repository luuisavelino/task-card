package request

type CardRequest struct {
	Title      string `json:"title" biding:"nonzero"`
	Summary    string `json:"summary" biding:"max=2500"`
	DueDate    string `json:"due_date"`
	CardStatus string `json:"card_status" biding:"nonzero"`
	UserId     int    `json:"user_id" biding:"nonzero"`
}
