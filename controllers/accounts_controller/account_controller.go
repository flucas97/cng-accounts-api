package accounts_controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/services/accounts_service"
	repository_service "github.com/flucas97/CNG-checknogreen/account/services/repository_service/cannabis"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
	"github.com/flucas97/CNG-checknogreen/account/utils/success_response"
	"github.com/gin-gonic/gin"
)

var (
	accountsService    = accounts_service.AccountsService
	cRepositoryService = repository_service.CannabisRepositoryService
)

// New handles the request from the server and persist the new account into psql
func New(c *gin.Context) {
	var account accounts.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		RestErr := error_factory.NewBadRequestError("Invalid JSON body")
		c.JSON(RestErr.Status, RestErr)
		fmt.Println("caiuuuuuuuu", err)
		return
	}

	result, err := accountsService.New(account)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	cannabisRepositoryID, err := cRepositoryService.NewRepository(result)
	if err != nil {
		fmt.Println(err)
		c.JSON(err.Status, err)
		return
	}
	/*
		TODO: save cannabis repository ID into psql
	*/

	fmt.Println("repository ID from account ", cannabisRepositoryID)
	c.Header("Repository-id", cannabisRepositoryID)
	c.Header("Account-id", strconv.Itoa(int(result.ID)))
	c.Header("Nick-name", account.Name)
	c.JSON(http.StatusCreated, result)
}

// Validate handles the request to check if the passed credentials are valid against psql
func Validate(c *gin.Context) {
	var account accounts.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		RestErr := error_factory.NewBadRequestError("invalid JSON body")
		c.JSON(RestErr.Status, RestErr)
		return
	}
	result, err := accountsService.Validate(account)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if !result {
		c.JSON(http.StatusNotFound, result)
		return
	}
	c.Header("Nick-name", account.Name)
	c.JSON(http.StatusOK, success_response.Found("successfully validate"))

}

// ShowDetails get all info about one specific account
func ShowDetails(c *gin.Context) {
	var accountName accounts.Account

	if err := c.ShouldBindJSON(&accountName); err != nil {
		RestErr := error_factory.NewBadRequestError("invalid JSON body")
		c.JSON(RestErr.Status, RestErr)
		return
	}

	account, err := accountsService.ShowDetails(accountName)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, account.Marshall())
}
