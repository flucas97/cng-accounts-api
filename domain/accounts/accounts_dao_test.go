package accounts

import (
	"fmt"
	"testing"

	"github.com/flucas97/CNG-checknogreen/account/utils/crypto"
	"github.com/flucas97/CNG-checknogreen/account/utils/date"
)

type accountFixture struct {
	testName     string
	expectError  bool
	errorMessage string

	mockAccount *Account
}

func (a *accountFixture) setup(t *testing.T) {
	a.mockAccount = newAccountMock()
}

func TestCreate(t *testing.T) {
	table := []accountFixture{
		{
			testName:    "Create new account with all valid parameters",
			expectError: false,
		},
		{
			testName:    "Invalid parameters",
			expectError: true,
		},
	}

	for index, f := range table {
		// t.Run is subtesting
		t.Run(fmt.Sprintf("%v-%v", index, f.testName), func(t *testing.T) {
			f.setup(t)

			err := f.mockAccount.Create()
			if !f.expectError && err != nil {
				t.Fatal("an error was not expected, but occurred")
			}
		},
		)
	}
}

func newAccountMock() *Account {
	return &Account{
		Name:      "abc",
		Email:     "abc@def.com",
		Status:    "active",
		Language:  "PT-br",
		Password:  crypto.GetMd5("abc123"),
		CreatedAt: date.GetNowString(),
		UpdatedAt: date.GetNowString(),
	}
}
