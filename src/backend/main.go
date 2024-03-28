package main

import (
	"database/sql"
	"fmt"
	"itacademy/gamble"
	"itacademy/raffle"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://admin:123@db:5432/itacademy?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Retry connecting to the database with exponential backoff
	retries := 0
	maxRetries := 5
	for db == nil {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			retries++
			if retries >= maxRetries {
				fmt.Println("failed to connect to the database after maximum retries")
				return
			}
			fmt.Printf("Failed to connect to the database (attempt %d/%d). Retrying in 5 seconds...\n", retries, maxRetries)
			time.Sleep(5 * time.Second)
		}
	}

	// Check if the database is ready by pinging it
	for {
		if err := db.Ping(); err != nil {
			fmt.Println("Error pinging the database. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	gambleHandler := gamble.NewHandler(db)
	raffleHandler := raffle.NewHandler(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.POST("/gamble/create", gambleHandler.Create)
	e.GET("/gamble/list", gambleHandler.List)
	e.POST("/raffle/create", raffleHandler.CreateRaffle)
	e.POST("/raffle/generate", raffleHandler.GenerateRaffle)
	e.POST("/raffle/start", raffleHandler.StartRaffle)
	e.POST("/raffle/close", raffleHandler.CloseRaffle)
	e.GET("/raffle/last", raffleHandler.LastRaffle)
	e.GET("/raffle/stats", raffleHandler.WinnersStats)
	e.GET("/raffle/count", raffleHandler.AllGambleNumbers)
	e.Logger.Fatal(e.Start(":8080"))
}
