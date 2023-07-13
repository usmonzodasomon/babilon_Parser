package db

import (
	"database/sql"
	"fmt"

	"github.com/usmonzodasomon/babilon_parser/utils"
)

var DB *sql.DB

func InitDB() *sql.DB {
	settings := utils.AppSettings.PostgresSettings
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		settings.User, settings.Password, settings.Database,
		settings.Server, settings.Port, settings.SSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func StartDBConnection() {
	DB = InitDB()
	up()
}

func CloseDBConnection(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
