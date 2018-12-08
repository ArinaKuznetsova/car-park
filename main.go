package main

import (
	"github.com/DonRIn/carpark/db"
	"github.com/DonRIn/carpark/routes"
	"log"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	routes.Run()
}
