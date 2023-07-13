package utils

import (
	"flag"
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

	getFlags()
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
	accountID := flag.Int("account_id", -1, "Account ID")
	tclass := flag.Int("tclass", -1, "TClass")
	sourceIP := flag.String("source_ip", "", "Source IP")
	destinationIP := flag.String("destination_ip", "", "Destination IP")
	flag.Parse()
	AppSettings.Flags.AccountID = int64(*accountID)
	AppSettings.Flags.Tclass = int64(*tclass)
	AppSettings.Flags.SourceIP = *sourceIP
	AppSettings.Flags.DestinationIP = *destinationIP
}

func SaveToFile(file *os.File, data string) error {
	_, err := file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
