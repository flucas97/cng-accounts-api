package accounts

// Account model
type Account struct {
	ID                   int64    `json:"id"`
	Name                 string   `json:"name"`
	CannabisRepositoryID int64    `belongs_to:"cannabis_repository_id" json:"cannabis_repository_id"`
	Status               string   `json:"status"`
	AvaliableFeatures    []string `json:"avaliable_features"`
	Language             string   `json:"language"`
	Password             string   `json:"password"`
	Token                string   `json:"token"`
	Country              string   `json:"country"`
	City                 string   `json:"city"`
	Description          string   `json:"description"`
	CreatedAt            string   `json:"created_at`
	UpdatedAt            string   `json:"update_at`
	Email                string   `json:"email"`
}
