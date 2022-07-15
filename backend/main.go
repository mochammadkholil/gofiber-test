package main

import (
	"backendsvc/api"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("fiber_db")
	if connStr == "" {
		log.Fatal("please set env variable name fiber_db with this format, host=%s port=%d user=%s password=%s dbname=%s sslmode=disable")
	}
	//connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "bismillah", "fiberdb")
	//connStr := "postgresql://postgres:bismillah@localhost/fiberdb?sslmode=disable"
	// Connect to database

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := api.Route(fiber.New(), db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "2222"
	}
	app.Static("/", "./public")

	log.Fatalln(app.Listen(fmt.Sprintf("localhost:%v", port)))
}
