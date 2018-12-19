package main

import (
	"github.com/ArinaKuznetsova/car-park/db"
	"github.com/ArinaKuznetsova/car-park/routes"
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
