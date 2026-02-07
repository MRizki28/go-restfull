package routes

import (
	"go-restfull/app/controller"
	"go-restfull/app/repository"
	"go-restfull/app/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	ProductController := controller.NewProductController(productService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api/v1")
	api.Get("/product", ProductController.GetAllProducts)
	api.Post("/product/create", ProductController.CreateDataProduct)

}
