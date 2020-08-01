package accounts_service

type AccountsServiceInterface interface {
	Create()
	Update()
	Login()
	Validate()
	Delete()
	Freeze()
}

type accountsService struct{}

var (
	// AccountsService interface for other services
	AccountsService AccountsServiceInterface = &accountsService{}
)

// Create Account
func (as *accountsService) Create() {

}

// Update Account
func (as *accountsService) Update() {

}

// Login Account
func (as *accountsService) Login() {

}

// Validate Account
func (as *accountsService) Validate() {

}

// Delete Account
func (as *accountsService) Delete() {

}

// Freeze Account
func (as *accountsService) Freeze() {

}
