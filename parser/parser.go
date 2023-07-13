package parser

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"sync"

	"github.com/usmonzodasomon/babilon_parser/db"
	"github.com/usmonzodasomon/babilon_parser/models"
)

var batchSize = 120 // Размер пакета для пакетной вставки

func ParseBinaryData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(175, 0)
	if err != nil {
		return err
	}

	dataChan := make(chan []models.DBData) // Канал для передачи буферов записей
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer close(dataChan)
		defer wg2.Done()
		buffer := make([]models.DBData, 0, batchSize) // Буфер для сбора записей
		for {
			data := models.ParserData{}

			err = binary.Read(file, binary.LittleEndian, &data)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			dbdata := Decode(&data)
			buffer = append(buffer, dbdata)

			// Если буфер достиг размера пакета, отправляем его в канал
			if len(buffer) == batchSize {
				dataChan <- buffer
				buffer = make([]models.DBData, 0, batchSize) // Очищаем буфер
			}
		}

		// Отправляем оставшиеся записи в канал
		if len(buffer) > 0 {
			dataChan <- buffer
		}
	}()

	var NumGor = runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(NumGor)
	for i := 0; i < NumGor; i++ {
		go func() {
			defer wg.Done()
			dbgor := db.InitDB()
			defer db.CloseDBConnection(dbgor)
			for buffer := range dataChan {
				if err := db.SaveData(dbgor, buffer); err != nil {
					panic(err)
				}
			}
		}()
	}

	fmt.Println("writing")
	wg2.Wait()
	wg.Wait()
	return nil
}

func Decode(data *models.ParserData) models.DBData {
	// Преобразование и добавление записи в буфер
	sourceIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(sourceIP, data.SourceIP)
	destinationIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(destinationIP, data.DestinationIP)
	NfSourceIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(NfSourceIP, data.NfSourceIP)

	dbdate := models.DBData{
		SourceIP:      sourceIP.String(),
		DestinationIP: destinationIP.String(),
		Packets:       data.Packets,
		Bytes:         data.Bytes,
		Sport:         data.Sport,
		Dport:         data.Dport,
		Proto:         data.Proto,
		AccountID:     data.AccountID,
		Tclass:        data.Tclass,
		DateTime:      data.DateTime,
		NfSourceIP:    NfSourceIP.String(),
	}
	return dbdate
}
