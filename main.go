package main

import (
	"fmt"
	"github.com/sheshankkodam/go-postgres/handlers"
	"log"
	"net/http"
	"github.com/sheshankkodam/go-postgres/db"
)

func main() {
	log.Println("Hello go-postgres")
	http.HandleFunc("/home", handlers.HelloHandler)

	go initDb()

	fmt.Println("Listening on port :3333")
	http.ListenAndServe(":3333", nil)
}

func initDb()  {
	log.Println("Initiating database")
	p := db.PostgresService{}
	if err := p.Insert("sheshank", "SSE", "05-08-2018"); err != nil {
		fmt.Printf("Error inserting, err=%v", err)
	}
}
