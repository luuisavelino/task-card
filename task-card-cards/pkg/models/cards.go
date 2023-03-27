package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/luuisavelino/task-card-cards/pkg/database"
)

type Card struct {
	Id         int    `json:"id"`
	Title      string `json:"title" binding:"required"`
	Summary    string `json:"summary"`
	DueDate    string `json:"due_date"`
	CardStatus string `json:"card_status" binding:"required"`
	UserId     int    `json:"user_id"`
}

func GetCards() ([]Card, error) {
	db := database.ConnectsWithDatabase()

	selectCards, err := db.Query("select * from cards order by id asc")
	if err != nil {
		log.Println(err)
		return nil, errors.New("error returning values")
	}

	c := Card{}
	cards := []Card{}

	for selectCards.Next() {
		var id, userId int
		var title, summary, cardStatus, dueDate string
		var err = selectCards.Scan(&id, &title, &summary, &dueDate, &cardStatus, &userId)
		if err != nil {
			log.Println(err)
			return nil, errors.New("error returning values")
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
		return Card{}, errors.New("error returning values")
	}

	card := Card{}
	for selectCard.Next() {
		var id, userId int
		var title, summary, cardStatus, dueDate string

		err = selectCard.Scan(&id, &title, &summary, &dueDate, &cardStatus, &userId)
		if err != nil {
			log.Println(err)
			return Card{}, errors.New("error returning values")
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

func CreateNewCard(card Card) error {
	db := database.ConnectsWithDatabase()

	insertCardIntoDatabase, err := db.Prepare("insert into cards(title, summary, due_date, card_status, user_id) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return errors.New("error returning values")
	}

	insertCardIntoDatabase.Exec(card.Title, card.Summary, card.DueDate, card.CardStatus, card.UserId)

	defer db.Close()
	return nil
}

func DeleteCard(cardId, userId int) error {
	db := database.ConnectsWithDatabase()

	executeIsManager, err := db.Query("select True from users join roles on roles.id = users.role_id where users.id = " + fmt.Sprint(userId) + " and roles.role_  = 'manager'")
	if err != nil {
		log.Println(err)
		return errors.New("error when deleting card")
	}

	var isManager bool
	for executeIsManager.Next() {
		executeIsManager.Scan(&isManager)
	}

	if isManager {
		deleteCard, err := db.Prepare("delete from cards where id=?")
		if err != nil {
			log.Println(err)
			return errors.New("error when deleting card")
		}

		deleteCard.Exec(cardId)
		defer db.Close()
		return nil
	}

	return errors.New("no permission to delete card")
}

func UpdateCard(card Card) error {
	db := database.ConnectsWithDatabase()

	updateCard, err := db.Prepare("Update cards set title=?, summary=?, card_status=?, due_date=? where id=?")
	if err != nil {
		log.Println(err)
		return errors.New("error returning values")
	}

	updateCard.Exec(card.Title, card.Summary, card.CardStatus, card.DueDate, card.Id)
	defer db.Close()
	return nil
}
