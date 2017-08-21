package main

import (
	"log"

	"github.com/loongy/jaguar/storage/db"
)

func clean(db *db.DB) error {
	log.Println("cleaning...")
	if err := cleanUserTable(db); err != nil {
		return err
	}
	log.Println("cleaning done.")
	return nil
}

func cleanUserTable(db *db.DB) error {
	log.Print("cleaning user table... ")
	_, err := db.Exec(`
		DELETE FROM users`)
	if err != nil {
		return err
	}
	log.Println("done.")
	return nil
}