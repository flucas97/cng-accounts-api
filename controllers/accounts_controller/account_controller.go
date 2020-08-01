package accounts_controller

import (
	"net/http"

	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/services/accounts_service"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
	"github.com/gin-gonic/gin"
)

var (
	accountsService = accounts_service.AccountsService
)

func Create(c *gin.Context) {
	var account accounts.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		RestErr := error_factory.NewBadRequestError("Invalid JSON body")
		c.JSON(RestErr.Status, RestErr)
		return
	}

	result, err := accountsService.Create(account)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
