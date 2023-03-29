package models

type CardDomainInterface interface {
	GetTitle() string
	GetSummary() string
	GetDueDate() string
	GetCardStatus() string
	GetUserId() int
}

func NewCardDomain(title, summary, dueDate, cardStatus string, userId int) CardDomainInterface {
	return &cardDomain{
		title:      title,
		summary:    summary,
		dueDate:    dueDate,
		cardStatus: cardStatus,
		userId:     userId,
	}
}

type cardDomain struct {
	title      string
	summary    string
	dueDate    string
	cardStatus string
	userId     int
}

func (u *cardDomain) GetTitle() string {
	return u.title
}
func (u *cardDomain) GetSummary() string {
	return u.summary
}
func (u *cardDomain) GetDueDate() string {
	return u.dueDate
}
func (u *cardDomain) GetCardStatus() string {
	return u.cardStatus
}
func (u *cardDomain) GetUserId() int {
	return u.userId
}
