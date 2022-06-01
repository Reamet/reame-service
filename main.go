package main

import (
	"bsc-scan-data-service/database"
	"bsc-scan-data-service/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main () {
	app :=  fiber.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	database.InitDbConfig()


	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	router.SetRouter(app)
	log.Fatal(app.Listen(":8082"))

}