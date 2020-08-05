package accounts_db

import (
	"database/sql"
	"fmt"

	"github.com/flucas97/CNG-checknogreen/account/utils/logger"
	_ "github.com/lib/pq"
)

const (
	psql_cng_accounts_username = "psql_cng_accounts_username"
	psql_cng_accounts_password = "psql_cng_accounts_password"
	psql_cng_accounts_root     = "psql_cng_accounts_root"
	psql_cng_accounts_schema   = "psql_cng_accounts_schema"
	psql_cng_accounts_port     = 5432
)

// need to use env instead hardcoded
var (
	Client   *sql.DB
	username = "postgres"
	password = "postgres"
	root     = "db"
	schema   = "cng_account"
	port     = psql_cng_accounts_port
)

func init() {
	var err error

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		root,
		port,
		schema,
	)

	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("cannot open database connection", err)
		panic(err)
	}

	logger.Info(username)

	err = Client.Ping()
	if err != nil {
		logger.Error("cannot ping database", err)
		panic(err)
	}

	logger.Info("accounts_db successfuly connected")
}
