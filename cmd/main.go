package main

import (
	"fmt"
	"log"
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

	filename := "ip60.utm"
	if err := parser.ParseBinaryData(filename); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(time.Since(start))
}
