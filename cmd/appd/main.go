package main

import (
	"database/sql"
	"github.com/adrianbrad/ddd-layout/internal/domain/inmem"
	"github.com/adrianbrad/ddd-layout/internal/domain/psql"
	"github.com/adrianbrad/ddd-layout/internal/http"
	"log"
)

func main() {
	// Connect to database.
	db, err := sql.Open("psql", "connection info...")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create services.
	us := psql.NewUserService(db)

	pubSub := inmem.NewPubSub()

	_ = http.NewServer(pubSub, us)

	// start http server...
}
