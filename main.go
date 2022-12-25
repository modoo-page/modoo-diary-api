package main

import (
	"golang-5252/api"
	"golang-5252/database"
	"golang-5252/server"
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.SetupDatabase()
	app := server.Create()
	api.Route(app)
	app.Listen(":5252")
}
