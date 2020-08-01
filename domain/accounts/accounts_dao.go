package accounts

import (
	"fmt"
	"strings"

	"github.com/flucas97/CNG-checknogreen/account/db/postgres/accounts_db"
	"github.com/flucas97/CNG-checknogreen/account/utils/crypto"
	"github.com/flucas97/CNG-checknogreen/account/utils/date"

	"github.com/flucas97/CNG-checknogreen/account/utils/logger"
	"github.com/flucas97/bookstore/users-api/utils/error_factory"
)

const (
	queryCreateAccount = ("INSERT INTO accounts(name, language, password, country, city, description) VALUES (?, ?, ?, ?, ?, ?);")
	statusActive       = "active"
	statusEnded        = "ended"
)

// Create Account
func (account *Account) Create() *error_factory.restErr {
	stmt, err := accounts_db.Client.Prepare(queryCreateAccount)
	if err != nil {
		logger.Error("error while preparing Create query", err)
		return errors.NewInternalServerError("an error occurred while creating account. Try again")
	}
	defer stmt.Close()

	account.CreatedAt, account.UpdatedAt, account.Status = date.GetNowString(), date.GetNowString(), statusActive
	account.Password = crypto.GetMd5(account.Password)

	insertResult, err := stmt.Exec(account.Name, account.Language, account.Password, account.Country, account.City, account.Description, account.Status, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return error_factory.NewBadRequestError(fmt.Sprintf("account '%v' already exists", account.Name))
		}
	}
	accountID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error while saving account in database query", err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save account: %s", err.Error()))
	}

	account.ID = accountID
	return nil
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
