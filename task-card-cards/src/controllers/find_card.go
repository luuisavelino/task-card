package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	_ "github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/configuration/validation"
	"github.com/luuisavelino/task-card-cards/src/controllers/model/request"
	"github.com/luuisavelino/task-card-cards/src/models"
	"github.com/luuisavelino/task-card-cards/src/view"
	"go.uber.org/zap"
)

//	@BasePath	/api/v1

// FindCards godoc
//	@Summary		Get all cards
//	@Description	Get all cards
//	@Tags			cards
//	@Accept			json
//	@Produce		json
//	@Param			action	body		request.ActionRequest	true	"Action info"
//	@Success		200		{object}	[]response.CardResponse
//	@Failure		400		{object}	rest_err.RestErr
//	@Router			/cards [get]
func (cc *cardControllerInterface) FindCards(c *gin.Context) {
	logger.Info("Init FindCards controller",
		zap.String("journey", "findCards"),
	)

	var actionRequest request.ActionRequest

	if errValidate := c.ShouldBindJSON(&actionRequest); errValidate != nil {
		logger.Error("Error trying to validate card info", errValidate)
		resterr := validation.ValidateCardError(errValidate)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	actionDomain := models.NewActionDomain(
		actionRequest.UserId,
	)

	domain, err := cc.service.FindCards(actionDomain)
	if err != nil {
		logger.Error("Error when trying to find user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindCards controller executed successfully",
		zap.String("journey", "findCards"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

//	@BasePath	/api/v1

// FindCardById godoc
//	@Summary		Get a card
//	@Description	Get a card
//	@Tags			cards
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Card id"
//	@Param			action	body		request.ActionRequest	true	"Action info"
//	@Success		200		{object}	[]response.CardResponse
//	@Failure		400		{object}	rest_err.RestErr
//	@Router			/cards/{id} [get]
func (cc *cardControllerInterface) FindCardById(c *gin.Context) {
	logger.Info("Init FindCardById controller",
		zap.String("journey", "findCardById"),
	)

	cardId, errValidate := strconv.Atoi(c.Params.ByName("id"))
	if errValidate != nil {
		logger.Error("Error trying to move card id", errValidate)
		resterr := validation.ValidateCardError(errValidate)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	var actionRequest request.ActionRequest

	if errValidate := c.ShouldBindJSON(&actionRequest); errValidate != nil {
		logger.Error("Error trying to validate card info", errValidate)
		resterr := validation.ValidateCardError(errValidate)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	actionDomain := models.NewActionDomain(
		actionRequest.UserId,
	)

	domain, err := cc.service.FindCardById(cardId, actionDomain)
	if err != nil {
		logger.Error("Error when trying to find user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindCardById controller executed successfully",
		zap.String("journey", "findCardById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
