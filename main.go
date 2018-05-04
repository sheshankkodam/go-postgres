package main

import (
	"log"
	"net/http"
	"github.com/sheshankkodam/go-postgres/handlers"
	"fmt"
)

func main()  {
	log.Println("Hello go-postgres")
	http.HandleFunc("/home", handlers.HelloHandler)
	fmt.Println("Listening on port :3333")
	http.ListenAndServe(":3333", nil)
}