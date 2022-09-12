package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"sgavrylenko/pg-intro/app"
	"sgavrylenko/pg-intro/website"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:Ohfaeyi7Zahz@localhost:5432/website")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	websiteRepository := website.NewPostgreSQLClassicRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}
