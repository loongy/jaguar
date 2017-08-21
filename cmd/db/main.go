package main

import (
	"flag"
	"log"
	"os"

	"github.com/loongy/jaguar/storage/db"
)

func main() {
	// Parse flags
	dbURL := flag.String("url", "", "database connection url")
	flag.Parse()

	// Open database connection and ping
	db, err := db.NewPostgresDB(*dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Which sub command
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "reset":
			if err := reset(db); err != nil {
				log.Fatal(err)
			}
		case "setup":
			if err := setup(db); err != nil {
				log.Fatal(err)
			}
		case "seed":
			if err := seed(db); err != nil {
				log.Fatal(err)
			}
		case "clean":
			if err := clean(db); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("Unknown subcommand '%v'", os.Args[1])
		}
	}
}
