package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"sgavrylenko/pg-intro/app"
	"sgavrylenko/pg-intro/website"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://postgres:Ohfaeyi7Zahz@localhost:5432/website")
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	websiteRepository := website.NewPostgreSQLPGXRepository(dbpool)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}
