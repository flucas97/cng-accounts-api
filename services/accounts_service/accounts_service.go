package accounts_service

import (
	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
)

const (
	statusActive = "active"
	statusEnded  = "ended"
	statusFreeze = "freeze"
)

type AccountsServiceInterface interface {
	Create(accounts.Account) (*accounts.Account, *error_factory.RestErr)
	Update()
	Validate(accounts.Account) (bool, *error_factory.RestErr)
	Delete()
	Freeze()
	ShowDetails(accounts.Account) (*accounts.Account, *error_factory.RestErr)
}

type accountsService struct{}

var (
	// AccountsService interface for other services
	AccountsService AccountsServiceInterface = &accountsService{}
)

// Create Account
func (as *accountsService) Create(account accounts.Account) (*accounts.Account, *error_factory.RestErr) {
	if err := account.Validate(); err != nil {
		return nil, err
	}

	account.PrepareFields(statusActive)

	if err := account.Create(); err != nil {
		return nil, err
	}

	return &account, nil
}

// Update Account
func (as *accountsService) Update() {

}

// Validate Account
func (as *accountsService) Validate(credentials accounts.Account) (bool, *error_factory.RestErr) {
	credentials.EncryptPassword()

	requestLogin := &accounts.Account{
		Name:     credentials.Name,
		Password: credentials.Password,
	}

	result, err := requestLogin.ValidateAccount()
	if err != nil {
		return result, err
	}

	return result, nil
}

// Delete Account
func (as *accountsService) Delete() {

}

// Freeze Account
func (as *accountsService) Freeze() {

}

func (as *accountsService) ShowDetails(account accounts.Account) (*accounts.Account, *error_factory.RestErr) {
	if err := account.ShowDetails(); err != nil {
		return nil, err
	}

	return &account, nil
}
