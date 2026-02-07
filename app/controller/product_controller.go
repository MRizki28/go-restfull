package controller

import (
	"go-restfull/app/request"
	"go-restfull/app/service"
	"go-restfull/app/trait"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ProductController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) GetAllProducts(c *fiber.Ctx) error {
	query := c.Query("search")
	products, err := pc.service.GetAllProducts(query)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to retrieve products",
			"error":   err.Error(),
		})
	}

	if len(products) == 0 {
		return c.Status(404).JSON(trait.DataNotFoundResponse("No products found"))
	}

	return c.Status(200).JSON(trait.SuccessResponse("Products retrieved successfully", products))
}

func (pc *ProductController) CreateDataProduct(c *fiber.Ctx) error {
	req := new(request.CreateProductRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(trait.ErrorResponse("Bad Request", "Invalid request body", err.Error()))
	}
	if err := validate.Struct(req); err != nil {
		errorss := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errorss[e.Field()] =  e.Tag() 
		}
		return c.Status(422).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "Validation failed",
			"errors":  errorss,
		})
	}

	product, err := pc.service.CreateDataProduct(*req)
	if err != nil {
		return c.Status(500).JSON(trait.ErrorResponse("Internal Server Error", "Failed to create product", err.Error()))
	}

	return c.Status(201).JSON(trait.SuccessResponse("Product created successfully", product)) 
}

func (pc *ProductController) GetDataById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(trait.ErrorResponse("Bad Request", "Invalid product ID", err.Error()))
	}

	product, err := pc.service.GetDataById(id)
	if err != nil {
		return c.Status(404).JSON(trait.DataNotFoundResponse("Product not found"))
	}

	return c.Status(200).JSON(trait.SuccessResponse("Product retrieved successfully", product))
}

func (pc *ProductController) UpdateData(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(
			trait.ErrorResponse("Bad Request", "Invalid product ID", err.Error()),
		)
	}

	req := new(request.CreateProductRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(
			trait.ErrorResponse("Bad Request", "Invalid request body", err.Error()),
		)
	}

	if err := validate.Struct(req); err != nil {
		errorss := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errorss[e.Field()] = e.Tag()
		}
		return c.Status(422).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "Validation failed",
			"errors":  errorss,
		})
	}

	product, err := pc.service.UpdateData(id, *req)
	if err != nil {

		errMsg := err.Error()

		if strings.Contains(errMsg, "record not found") {
			return c.Status(404).JSON(
				trait.DataNotFoundResponse("Product not found"),
			)
		}

		if strings.Contains(errMsg, "no data changed") {
			return c.Status(400).JSON(
				trait.ErrorResponse(
					"Bad Request",
					"No data was changed",
					"Request data is identical to existing data",
				),
			)
		}

		return c.Status(500).JSON(
			trait.ErrorResponse(
				"Internal Server Error",
				"Failed to update product",
				errMsg,
			),
		)
	}

	return c.Status(200).JSON(
		trait.SuccessResponse("Product updated successfully", product),
	)
}

func (pc *ProductController) DeleteData(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(trait.ErrorResponse("Bad Request", "Invalid product ID", err.Error()))
	}
	err = pc.service.DeleteData(id)

	if err != nil {
		errMsg := err.Error()

		if strings.Contains(errMsg, "product not found") {
			return c.Status(404).JSON(trait.DataNotFoundResponse("Product not found"))
		}

		return c.Status(500).JSON(trait.ErrorResponse("Internal Server Error", "Failed to delete product", errMsg))
	}

	return c.Status(200).JSON(trait.DeleteDataSuccessResponse("Product deleted successfully"))
}

