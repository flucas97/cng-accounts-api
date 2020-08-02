package accounts

type AccountDetails struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (account Account) Marshall() interface{} {
	return AccountDetails{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		City:      account.City,
		State:     account.State,
		Country:   account.Country,
		Language:  account.Language,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
