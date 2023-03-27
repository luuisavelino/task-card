package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/pkg/models"
)

type requestReturn struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Users(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func User(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", "invalid id",
		})
		return
	}

	user, err := models.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", "invalid id",
		})
		return
	}

	if err = models.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, requestReturn{
		"success", "user deleted",
	})
}

func CreateUser(c *gin.Context) {
	var user models.User

	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		switch key {
		case "username":
			user.Username = value[0]
		case "userpass":
			user.Userpass = value[0]
		case "role_id":
			roleId, err := strconv.Atoi(value[0])
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, requestReturn{
					"error", "invalid role id",
				})
				return
			}
			if roleId != 1 && roleId != 2 {
				log.Println("invalid role id")
				c.JSON(http.StatusBadRequest, requestReturn{
					"error", "invalid role id",
				})
				return
			}
			user.RoleId = roleId
		}
	}

	if err := models.CreateNewUser(user); err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, requestReturn{
		"success", "user created",
	})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", "invalid id",
		})
		return
	}

	var user models.User

	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		switch key {
		case "username":
			user.Username = value[0]
		case "userpass":
			user.Userpass = value[0]
		case "role_id":
			roleId, err := strconv.Atoi(value[0])
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, requestReturn{
					"error", "invalid role id",
				})
				return
			}
			if roleId != 1 && roleId != 2 {
				c.JSON(http.StatusBadRequest, requestReturn{
					"error", "invalid role id",
				})
				return
			}
			user.RoleId = roleId
		}
	}

	user.Id = id
	models.UpdateUser(user)

	c.JSON(http.StatusOK, requestReturn{
		"success", "user updated",
	})
}
