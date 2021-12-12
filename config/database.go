package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	// init database
	dbHost := DATABASE_HOST
	dbPort := DATABASE_PORT
	dbUser := DATABASE_USERNAME
	dbPass := DATABASE_PASSWORD
	dbName := DATABASE_NAME
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}
	return dbConn
}
