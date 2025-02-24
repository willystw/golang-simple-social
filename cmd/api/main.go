package main

import (
	"log"

	"github.com/willystw/golang-simple-social/internal/db"
	"github.com/willystw/golang-simple-social/internal/env"
	"github.com/willystw/golang-simple-social/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:              env.GetString("DB_ADDR", "postgres://postgres:pwd@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns:      env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConnection: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:       env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConnection,
		cfg.db.maxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database conection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
