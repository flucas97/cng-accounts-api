package accounts

import (
	"strings"

	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
)

// Account model
type Account struct {
	ID                   int64    `json:"id"`
	Name                 string   `json:"name"`
	Email                string   `json:"email"`
	Password             string   `json:"password"`
	City                 string   `json:"city"`
	Country              string   `json:"country"`
	State                string   `json:"state"`
	AvaliableFeatures    []string `json:"avaliable_features"`
	CannabisRepositoryID int64    `belongs_to:"cannabis_repository_id" json:"cannabis_repository_id"`
	Description          string   `json:"description"`
	Language             string   `json:"language"`
	Status               string   `json:"status"`
	CreatedAt            string   `json:"created_at`
	UpdatedAt            string   `json:"updated_at`
}

func (a *Account) Validate() *error_factory.RestErr {
	a.Email = strings.TrimSpace(strings.ToLower(a.Email))
	a.Name = strings.TrimSpace(strings.ToLower(a.Email))

	if a.Email == "" {
		return error_factory.NewBadRequestError("Invalid email address")
	}

	if a.Name == "" {
		return error_factory.NewBadRequestError("Invalid name")
	}

	return nil
}
