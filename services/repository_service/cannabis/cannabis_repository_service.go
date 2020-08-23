package repository_service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
)

var (
	CannabisRepositoryService cannabisRepositoryInterface = &cannabisRepositoryService{}
	cannabisServiceURI                                    = os.Getenv("CANNABIS_BASE_URL")
)

type cannabisRepositoryInterface interface {
	NewRepository(*accounts.Account) (string, *error_factory.RestErr)
}

type cannabisRepositoryService struct{}

func (crs cannabisRepositoryService) NewRepository(a *accounts.Account) (string, *error_factory.RestErr) {
	values := map[string]interface{}{
		"name":       a.Name,
		"account_id": strconv.Itoa(int(a.ID)),
	}

	buffer := new(bytes.Buffer)

	json.NewEncoder(buffer).Encode(values)
	jsonValue, _ := json.Marshal(values)

	r, err := http.Post(cannabisServiceURI+"create-repository", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", error_factory.NewInternalServerError(err.Error())
	}

	return r.Header.Get("Repository-id"), nil
}
