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

// AccountsServiceInterface has all the methods implemented by this microservice
type AccountsServiceInterface interface {
	New(accounts.Account) (*accounts.Account, *error_factory.RestErr)
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

// New Account
func (as *accountsService) New(account accounts.Account) (*accounts.Account, *error_factory.RestErr) {
	if err := account.Validate(); err != nil {
		return nil, err
	}

	account.PrepareFields(statusActive)

	if err := account.New(); err != nil {
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
	/*
		TODO: implement Delete Account method
	*/
}

// Freeze Account
func (as *accountsService) Freeze() {
	/*
		TODO: implement Freeze Account method, with propose of inactivate an account
	*/
}

func (as *accountsService) ShowDetails(account accounts.Account) (*accounts.Account, *error_factory.RestErr) {
	if err := account.ShowDetails(); err != nil {
		return nil, err
	}

	return &account, nil
}
