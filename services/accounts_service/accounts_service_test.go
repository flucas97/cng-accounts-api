package accounts_service

import (
	"testing"

	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
)

type TestCase struct {
	credentials accounts.Account
	expected    bool
}

func TestValidate(t *testing.T) {
	cases := []TestCase{
		{
			credentials: accounts.Account{
				Name:     "abc",
				Password: "abc",
			},
			expected: false,
		},
	}

	for _, c := range cases {
		actualCredentials := c.credentials
		expected := c.expected

		result, _ := actualCredentials.ValidateAccount()

		if result != expected {
			t.Logf("Expected %v, got %v", expected, result)
			t.Fail()
		}
	}

}
