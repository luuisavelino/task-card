package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/luuisavelino/task-card-cards/pkg/database"
)

const (
	genericErrToUser = "error returning values"
)

type Card struct {
	Id         int    `json:"id"`
	Title      string `json:"title" validate:"nonzero"`
	Summary    string `json:"summary" validate:"max=2500"`
	DueDate    string `json:"due_date"`
	CardStatus string `json:"card_status" validate:"nonzero"`
	UserId     int    `json:"user_id" validate:"nonzero"`
}

func GetCards(cardId int) ([]Card, error) {
	db := database.ConnectsWithDatabase()

	selectUserRole, err := db.Query("select roles.role_ from users join roles on roles.id = users.role_id where users.id = " + fmt.Sprint(cardId))
	if err != nil {
		log.Println(err)
		return nil, errors.New(genericErrToUser)
	}

	var userRole string
	for selectUserRole.Next() {
		if err = selectUserRole.Scan(&userRole); err != nil {
			log.Println(err)
			return nil, errors.New(genericErrToUser)
		}
	}

	var queryToGetCards string
	switch userRole {
	case "manager":
		queryToGetCards = "select * from cards"
	case "technician":
		queryToGetCards = "select * from cards where user_id = " + fmt.Sprint(cardId)
	}

	selectCards, err := db.Query(queryToGetCards)
	if err != nil {
		log.Println(err)
		return nil, errors.New(genericErrToUser)
	}

	c := Card{}
	cards := []Card{}

	for selectCards.Next() {
		var id, userId int
		var title, summary, cardStatus, dueDate string
		if err = selectCards.Scan(&id, &title, &summary, &dueDate, &cardStatus, &userId); err != nil {
			log.Println(err)
			return nil, errors.New(genericErrToUser)
		}

		c.Id = id
		c.Title = title
		c.Summary = summary
		c.CardStatus = cardStatus
		c.DueDate = dueDate
		c.UserId = userId

		cards = append(cards, c)
	}

	defer db.Close()

	return cards, nil
}

func GetCard(id int) (Card, error) {
	db := database.ConnectsWithDatabase()

	selectCard, err := db.Query("select * from cards where id=?", id)
	if err != nil {
		log.Println(err)
		return Card{}, errors.New(genericErrToUser)
	}

	card := Card{}
	for selectCard.Next() {
		var id, userId int
		var title, summary, cardStatus, dueDate string

		err = selectCard.Scan(&id, &title, &summary, &dueDate, &cardStatus, &userId)
		if err != nil {
			log.Println(err)
			return Card{}, errors.New(genericErrToUser)
		}

		card.Id = id
		card.Title = title
		card.Summary = summary
		card.CardStatus = cardStatus
		card.DueDate = dueDate
		card.UserId = userId
	}

	defer db.Close()
	return card, nil
}
