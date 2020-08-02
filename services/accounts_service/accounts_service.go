package accounts_service

import (
	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/utils/crypto"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
)

type AccountsServiceInterface interface {
	Create(accounts.Account) (*accounts.Account, *error_factory.RestErr)
	Update()
	Login(accounts.Account) (bool, *error_factory.RestErr)
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
	if err := account.Validate(); err != nil {
		return nil, err
	}

	if err := account.Create(); err != nil {
		return nil, err
	}

	return &account, nil
}

// Update Account
func (as *accountsService) Update() {

}

// Login Account
func (as *accountsService) Login(credentials accounts.Account) (bool, *error_factory.RestErr) {
	requestLogin := &accounts.Account{
		Name:     credentials.Name,
		Password: crypto.GetMd5(credentials.Password),
	}

	result, err := requestLogin.Login()
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
