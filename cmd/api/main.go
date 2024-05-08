package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Uemerson/clean-go-api/internal/application/user/route"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		fmt.Println("Failed to connect to the database.")
		return
	}

	mux := http.NewServeMux()
	route.UserRoute(mux, db)
	fmt.Println("Server is running on port " + os.Getenv("API_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("API_PORT")), mux)
}
