package main

import (
	"log"
	"os"

	"github.com/loongy/jaguar-template/api"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening...")
	if err := server.ListenAndServe(":" + port); err != nil {
		log.Fatal(err)
	}
}
