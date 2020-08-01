package accounts_service

import (
	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
)

type AccountsServiceInterface interface {
	Create(account accounts.Account) (*accounts.Account, *error_factory.RestErr)
	Update()
	Login()
	Validate()
	Delete()
	Freeze()
}

type accountsService struct{}

var (
	// AccountsService interface for other services
	AccountsService AccountsServiceInterface = &accountsService{}
)

// Create Account
func (as *accountsService) Create(account accounts.Account) (*accounts.Account, *error_factory.RestErr) {
	if err := account.Create(); err != nil {
		return nil, err
	}

	return &account, nil
}

// Update Account
func (as *accountsService) Update() {

}

// Login Account
func (as *accountsService) Login() {

}

// Validate Account
func (as *accountsService) Validate() {

}

// Delete Account
func (as *accountsService) Delete() {

}

// Freeze Account
func (as *accountsService) Freeze() {

}
