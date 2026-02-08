package routes

import (
	"go-restfull/app/controller"
	"go-restfull/app/middleware"
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

	product := api.Group("/product" , middleware.AuthMiddleware)
	product.Get("/", ProductController.GetAllProducts)
	product.Post("/create", ProductController.CreateDataProduct)
	product.Get("/get/:id", ProductController.GetDataById)
	product.Post("/update/:id", ProductController.UpdateData)
	product.Delete("/delete/:id", ProductController.DeleteData)

}
