package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
	"vrs/utils"
)

var (
	username        = utils.Dotenv("DB_USERNAME")
	password        = utils.Dotenv("DB_PASSWORD")
	host            = utils.Dotenv("DB_HOST")
	port            = utils.Dotenv("DB_PORT")
	dbname          = utils.Dotenv("DB_DATABASE")
	applicationName = utils.Dotenv("APPLICATION_NAME")
)

func InitConnectionDatabase() {
	utils.DB = connectionBDPostgreSQL(username, password, host, port, dbname, applicationName)
	utils.DB.SetMaxOpenConns(25)
	utils.DB.SetMaxIdleConns(25)
	utils.DB.SetConnMaxLifetime(1 * time.Minute)
}

func connectionBDPostgreSQL(username, password, host, port, dbname, applicationName string) *sql.DB {
	dsn := "user=" + username + " password=" + password + " host=" + host + " port=" + port + " dbname=" + dbname + " application_name=" + applicationName + " sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		if err.Error() == "pq: database \""+dbname+"\" does not exist" {
			err := CreateDatabase(username, password, host, port, dbname)
			connectionBDPostgreSQL(username, password, host, port, dbname, applicationName)
			if err != nil {
				return nil
			}
		} else {
			panic(err.Error())
		}
	}
	return db
}
