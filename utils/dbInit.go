package utils

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DBInit initialize MySQL connection.
func DBInit() (*sql.DB, error) {
	logfile, er := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		panic(er.Error())
	}
	defer logfile.Close()

	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("DB_PORT")
	database := os.Getenv("DATABASE")
	dbDriver := os.Getenv("DRIVER_NAME")

	dataSourceName := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + database

	sql, err := sql.Open(dbDriver, dataSourceName)

	return sql, err
}
