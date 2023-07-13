package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/usmonzodasomon/babilon_parser/models"
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
