package accounts_db

import (
	"database/sql"
	"fmt"

	"github.com/flucas97/bookstore/users-api/logger"
	_ "github.com/lib/pq"
)

const (
	postgrsql_host           = "localhost"
	postgrsql_port           = 5432
	postgrsql_user           = "feijo"
	postgrsql_password       = "root"
	postgrsql_database_bname = "CNG-Account"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgrsql_host,
		postgrsql_port,
		postgrsql_user,
		postgrsql_password,
		postgrsql_database_bname,
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
