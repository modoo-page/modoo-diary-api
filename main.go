package main

import (
	"log"
	"math/rand"
	"modoo-diary-api/api"
	"modoo-diary-api/database"
	"modoo-diary-api/server"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())
	api.Route(app)
	app.Listen(":5252")
}
