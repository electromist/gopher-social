package main

import (
	"log"

	"github.com/electromist/gopher-social.git/internal/db"
	"github.com/electromist/gopher-social.git/internal/env"
	"github.com/electromist/gopher-social.git/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")

	// Transcript: "I just do 3 and 3 connections here... I don't want to be generous for these scripts"
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Store initialize
	store := store.PostgresStorage(conn)

	// internal/db package wala Seed function
	db.Seed(store, conn)
}
