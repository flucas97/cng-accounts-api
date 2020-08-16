package repository_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/flucas97/CNG-checknogreen/account/domain/accounts"
	"github.com/flucas97/CNG-checknogreen/account/utils/error_factory"
)

var (
	CannabisRepositoryService cannabisRepositoryInterface = &cannabisRepositoryService{}
	cannabisServiceURI                                    = "http://172.30.0.7:8083/api/"
)

type cannabisRepositoryInterface interface {
	NewRepository(*accounts.Account) (string, *error_factory.RestErr)
}

type cannabisRepositoryService struct{}

func (crs cannabisRepositoryService) NewRepository(a *accounts.Account) (string, *error_factory.RestErr) {
	fmt.Printf("account ID %v name %v \n", a.ID, a.Name)
	values := map[string]string{"name": a.Name, "account_id": strconv.Itoa(int(a.ID))}
	jsonValue, _ := json.Marshal(values)

	r, err := http.Post(cannabisServiceURI+"create-repository", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", error_factory.NewInternalServerError(err.Error())
	}

	return r.Header.Get("cannabis_repository_id"), nil
}
