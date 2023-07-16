package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/usmonzodasomon/babilon_parser/db"
	"github.com/usmonzodasomon/babilon_parser/parser"
	"github.com/usmonzodasomon/babilon_parser/utils"
)

func main() {
	start := time.Now()

	utils.ReadSettings()

	db.StartDBConnection()
	defer db.CloseDBConnection(db.DB)

	if flag.NFlag() == 0 {
		filename := "ipbig.utm"
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if err := parser.ParseBinaryData(file); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if err := db.GetDate(file); err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println(time.Since(start))
}
