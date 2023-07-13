package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/usmonzodasomon/babilon_parser/models"
	"github.com/usmonzodasomon/babilon_parser/utils"
)

func SaveData(db *sql.DB, data []models.DBData) error {
	var valueStrings []string
	var valueArgs []interface{}
	idx := 1
	for i := 0; i < len(data); i++ {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", idx, idx+1, idx+2, idx+3, idx+4, idx+5, idx+6, idx+7, idx+8, idx+9, idx+10))
		valueArgs = append(valueArgs, data[i].SourceIP, data[i].DestinationIP, data[i].Packets, data[i].Bytes, data[i].Sport, data[i].Dport, data[i].Proto, data[i].AccountID, data[i].Tclass, data[i].DateTime, data[i].NfSourceIP)
		idx += 11
	}
	_, err := db.Exec("INSERT INTO netflow (source_ip, destination_ip, packets, bytes, sport, dport, proto, account_id, tclass, date_time, nfsource_ip) VALUES "+strings.Join(valueStrings, ","), valueArgs...)
	if err != nil {
		return err
	}

	return nil
}

func GetDate(file *os.File) error {
	sqlQuery := "SELECT * FROM netflow WHERE 1=1"
	flags := utils.AppSettings.Flags
	if flags.AccountID != -1 {
		sqlQuery += fmt.Sprintf(" AND account_id = %d", flags.AccountID)
	}
	if flags.Tclass != -1 {
		sqlQuery += fmt.Sprintf(" AND tclass = %d", flags.Tclass)
	}
	if flags.SourceIP != "" {
		sqlQuery += fmt.Sprintf(" AND source_ip = '%s'", flags.SourceIP)
	}
	if flags.DestinationIP != "" {
		sqlQuery += fmt.Sprintf(" AND destination_ip = '%s'", flags.DestinationIP)
	}

	rows, err := DB.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var (
		sourceIP      string
		destinationIP string
		packets       int
		bytes         int
		sport         int
		dport         int
		proto         string
		accountID     int
		tclass        int
		dateTime      string
		nfSourceIP    string
	)
	for rows.Next() {
		// Чтение данных из результата запроса
		err = rows.Scan(&sourceIP, &destinationIP, &packets, &bytes, &sport, &dport, &proto, &accountID, &tclass, &dateTime, &nfSourceIP)
		if err != nil {
			panic(err)
		}

		// Вывод данных
		data := fmt.Sprintf("Source IP: %s, Destination IP: %s, Packets: %d, Bytes: %d, Sport: %d, Dport: %d, Proto: %s, Account ID: %d, TClass: %d, DateTime: %s, NF Source IP: %s\n",
			sourceIP, destinationIP, packets, bytes, sport, dport, proto, accountID, tclass, dateTime, nfSourceIP)
		if err := utils.SaveToFile(file, data); err != nil {
			return err
		}
	}
	return nil
}
