package main

import (
	"log"

	"github.com/willystw/golang-simple-social/internal/db"
	"github.com/willystw/golang-simple-social/internal/env"
	"github.com/willystw/golang-simple-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:pwd@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(
		addr,
		3,
		3,
		"15m")

	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
