package accounts_db

import (
	"database/sql"
	"fmt"
	"os"

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

var (
	Client   *sql.DB
	username = os.Getenv(psql_cng_accounts_username)
	password = os.Getenv(psql_cng_accounts_password)
	root     = os.Getenv(psql_cng_accounts_root)
	schema   = os.Getenv(psql_cng_accounts_schema)
	port     = psql_cng_accounts_port
)

func init() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		root,
		port,
		username,
		password,
		schema,
	)

	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("cannot open database connection", err)
		panic(err)
	}

	err = Client.Ping()
	if err != nil {
		logger.Error("cannot ping database", err)
		panic(err)
	}

	logger.Info("accounts_db successfuly connected")
}
