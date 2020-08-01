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

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgrsql_host,
		postgrsql_port,
		postgrsql_user,
		postgrsql_password,
		postgrsql_database_bname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("cannot open database connection", err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("cannot ping database", err)
		panic(err)
	}

	logger.Info("accounts_db successfuly connected")
}
