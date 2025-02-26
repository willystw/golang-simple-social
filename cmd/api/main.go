package main

import (
	"time"

	"github.com/willystw/golang-simple-social/internal/db"
	"github.com/willystw/golang-simple-social/internal/env"
	"github.com/willystw/golang-simple-social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			Gopher Social API
//	@description	API for Gopher Social, a social network of Gophers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:              env.GetString("DB_ADDR", "postgres://postgres:pwd@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns:      env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConnection: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:       env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3,
		},
	}
	//Logger

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConnection,
		cfg.db.maxIdleTime)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database conection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
