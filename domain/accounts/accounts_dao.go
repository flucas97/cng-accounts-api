package accounts

import (
	"strings"

	"github.com/flucas97/CNG-checknogreen/account/db/postgres/accounts_db"
	"github.com/flucas97/CNG-checknogreen/account/utils/crypto"
	"github.com/flucas97/CNG-checknogreen/account/utils/date"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
	"github.com/flucas97/CNG-checknogreen/account/utils/logger"
)

const (
	queryCreateAccount = ("INSERT INTO accounts (name, email, language, password, country, city, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;")
	queryLogin         = ("SELECT name, password FROM accounts WHERE name=$1 AND password=$2;")
	statusActive       = "active"
	statusEnded        = "ended"
)

// Create Account
func (account *Account) Create() *error_factory.RestErr {
	var err error
	// check if is everything OK accessing the DB
	if err = accounts_db.Client.Ping(); err != nil {
		panic(err)
	}

	account.CreatedAt, account.UpdatedAt, account.Status = date.GetNowString(), date.GetNowString(), statusActive
	account.Password = crypto.GetMd5(account.Password)

	accountID := 0
	err = accounts_db.Client.QueryRow(queryCreateAccount, account.Name, account.Email, account.Language, account.Password, account.Country, account.City, account.Description, account.Status, account.CreatedAt, account.UpdatedAt).Scan(&accountID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return error_factory.NewBadRequestError("email or account already exists.")
		}
		logger.Error("error querying row", err)
		return error_factory.NewInternalServerError("error saving account, try again")
	}

	account.ID = int64(accountID)
	return nil
}

// Update Account
func (account *Account) Update() {

}

// Login Account
func (credentials *Account) Login() (bool, *error_factory.RestErr) {
	var err error

	if err := accounts_db.Client.Ping(); err != nil {
		panic(err)
	}

	err = accounts_db.Client.QueryRow(queryLogin, credentials.Name, credentials.Password).Scan(&credentials.Name, &credentials.Password)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return false, error_factory.NewNotFoundError("not found")
		}

		logger.Error("error while preparing query", err)
		return false, error_factory.NewInternalServerError("error, try again")
	}
	return true, nil
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
