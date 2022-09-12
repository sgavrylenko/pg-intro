package main

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sgavrylenko/pg-intro/app"
	"sgavrylenko/pg-intro/website"
)

func main() {
	gormDB, err := gorm.Open(postgres.Open("postgres://postgres:Ohfaeyi7Zahz@localhost:5432/website"))
	if err != nil {
		log.Fatal(err)
	}

	websiteRepository := website.NewPostgreSQLGORMRepository(gormDB)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}
