package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/usmonzodasomon/babilon_parser/models"
)

var AppSettings models.AppSettings

func ReadSettings() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldn't open config file. Error is: ", err.Error())
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read `config.yaml` file. Error is: ", err.Error())
	}

	// getFlags()
	setupPostgres()
}

func setupPostgres() {
	AppSettings.PostgresSettings.Server = viper.GetString("db.host")
	AppSettings.PostgresSettings.Port = viper.GetString("db.port")
	AppSettings.PostgresSettings.User = viper.GetString("db.username")
	AppSettings.PostgresSettings.Database = viper.GetString("db.dbname")
	AppSettings.PostgresSettings.SSLMode = viper.GetString("db.sslmode")
	AppSettings.PostgresSettings.Password = os.Getenv("DB_PASSWORD")
}

func getFlags() {

}
