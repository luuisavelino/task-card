package models

import (
	"errors"
	"log"

	"github.com/luuisavelino/task-card-users/pkg/database"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Userpass string `json:"userpass" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleId   int    `json:"role_id"`
}

func GetAllUsers() ([]User, error) {

	db := database.ConnectsWithDatabase()

	selectAllUsers, err := db.Query("select * from users order by id asc")
	if err != nil {
		return nil, errors.New("error returning values")
	}

	u := User{}
	users := []User{}

	for selectAllUsers.Next() {
		var id, roleId int
		var username, userpass, email string

		err = selectAllUsers.Scan(&id, &username, &userpass, &email, &roleId)
		if err != nil {
			return nil, errors.New("error returning values")
		}

		u.Id = id
		u.Username = username
		u.Userpass = userpass
		u.Email = email
		u.RoleId = roleId

		users = append(users, u)
	}

	defer db.Close()

	return users, nil
}

func GetUser(id int) (User, error) {
	db := database.ConnectsWithDatabase()

	selectUser, err := db.Query("select * from users where id=?", id)
	if err != nil {
		log.Println(err)
		return User{}, errors.New("error returning values")
	}

	user := User{}
	for selectUser.Next() {
		var id, roleId int
		var username, userpass, email string

		err = selectUser.Scan(&id, &username, &userpass, &email, &roleId)
		if err != nil {
			log.Println(err)
			return User{}, errors.New("error returning values")
		}

		user.Id = id
		user.Username = username
		user.Userpass = userpass
		user.Email = email
		user.RoleId = roleId
	}

	defer db.Close()
	return user, nil
}

func CreateNewUser(user User) error {
	db := database.ConnectsWithDatabase()

	insertUserIntoDatabase, err := db.Prepare("insert into users(username, userpass, email, role_id) values(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return errors.New("error returning values")
	}

	insertUserIntoDatabase.Exec(user.Username, user.Userpass, user.Email, user.RoleId)
	defer db.Close()
	return nil
}

func DeleteUser(id int) error {
	db := database.ConnectsWithDatabase()

	deleteUser, err := db.Prepare("delete from users where id=?")
	if err != nil {
		log.Println(err)
		return errors.New("error returning values")
	}

	deleteUser.Exec(id)
	defer db.Close()
	return nil
}

func UpdateUser(user User) error {
	db := database.ConnectsWithDatabase()

	updateUser, err := db.Prepare("Update users set username=?, userpass=?, email=?, role_id=? where id=?")
	if err != nil {
		log.Println(err)
		return errors.New("error returning values")
	}

	updateUser.Exec(user.Username, user.Userpass, user.Email, user.RoleId, user.Id)
	defer db.Close()
	return nil
}
