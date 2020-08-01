package accounts

import (
	"github.com/flucas97/CNG-checknogreen/account/db/postgres/accounts_db"
	"github.com/flucas97/CNG-checknogreen/account/utils/errors_utils"
	"github.com/flucas97/CNG-checknogreen/account/utils/logger"
)

const (
	queryCreateAccount = ("INSERT INTO accounts(name, language, password, country, city, description) VALUES (?, ?, ?, ?, ?, ?);")
)

// Create Account
func (account *Account) Create() {
	stmt, err := accounts_db.Client.Prepare(queryCreateAccount)
	if err != nil {
		logger.Error("error while preparing Create query", err)
		return errors_utils.NewInternalServerError("an error occurred while creating account. Try again")
	}
	defer stmt.Close()

}

// Update Account
func (account *Account) Update() {

}

// Login Account
func (account *Account) Login() {

}

// Validate Account
func (account *Account) Validate() {

}

// Delete Account
func (account *Account) Delete() {

}

// Freeze Account
func (account *Account) Freeze() {

}
