package main

import (
	"github.com/DiasOrazbaev/FoodNear/database"
	"github.com/DiasOrazbaev/FoodNear/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	database.GetConnection()
	app := fiber.New(fiber.Config{
		AppName: "FoodNEAR",
		Prefork: true,
	})

	routes.InitRoutes(app)

	app.Use(logger.New())

	log.Fatalln(app.Listen(":3030"))
}
