package main

import (
	"fmt"
	"go-restfull/app/config"
	"go-restfull/app/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectionDatabase()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	app.Listen(":" + port)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
