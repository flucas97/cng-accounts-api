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
	queryCreateAccount = ("INSERT INTO account (name, email, language, password, country, city, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;")
	queryLogin         = ("SELECT name, password FROM account WHERE name=$1 AND password=$2;")
	queryShowDetails   = ("SELECT id, name, email, country, city, description, created_at, updated_at FROM account WHERE name=$1;")
	statusActive       = "active"
	statusEnded        = "ended"
)

// Create Account
func (account *Account) Create() *error_factory.RestErr {
	var err error
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
		return error_factory.NewInternalServerError("error while saving account, try again")
	}

	account.ID = int64(accountID)
	return nil
}

// Update Account
func (account *Account) Update() {

}

// ShowDetails get info about one account
func (account *Account) ShowDetails() *error_factory.RestErr {
	var err error

	if err := accounts_db.Client.Ping(); err != nil {
		panic(err)
	}

	err = accounts_db.Client.QueryRow(queryShowDetails, account.Name).Scan(&account.ID, &account.Name, &account.Email, &account.Country, &account.City, &account.Description, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return error_factory.NewNotFoundError("account not found")
		}

		logger.Error("error while preparing query", err)
		return error_factory.NewInternalServerError("error searching account informations, try again")
	}
	return nil

}

// ValidateAccount
func (credentials *Account) ValidateAccount() (bool, *error_factory.RestErr) {
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
		return false, error_factory.NewInternalServerError("error while requesting validate, try again")
	}
	return true, nil
}

// Delete Account
func (account *Account) Delete() {

}

// Freeze Account
func (account *Account) Freeze() {

}
