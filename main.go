package main

import (
	"bsc-scan-data-service/database"
	"bsc-scan-data-service/env"
	"bsc-scan-data-service/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	env.Load()
	database.InitDbConfig()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	router.SetRouter(app)

	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))

}
