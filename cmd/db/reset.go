package main

import (
	"log"

	"github.com/loongy/jaguar-template/storage/db"
)

func reset(db *db.DB) error {
	log.Println("resetting...")
	if err := resetUserTable(db); err != nil {
		return err
	}
	log.Println("resetting done.")
	return nil
}

func resetUserTable(db *db.DB) error {
	log.Print("resetting user table... ")
	_, err := db.Exec(`
		DROP TABLE IF EXISTS users`)
	if err != nil {
		return err
	}
	log.Println("done.")
	return nil
}
