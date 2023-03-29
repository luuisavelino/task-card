package models

type ActionDomainInterface interface {
	GetUserId() int
}

func NewActionDomain(userId int) ActionDomainInterface {
	return &cardDomain{
		userId: userId,
	}
}

type actionDomain struct {
	userId int
}

func (a *actionDomain) GetUserId() int {
	return a.userId
}
