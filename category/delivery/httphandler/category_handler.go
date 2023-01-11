package httphandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/middleware"
	"github.com/mrizalr/mini-project-evermos/model"
)

type categoryHandler struct {
	categoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(r fiber.Router, categoryUsecase domain.CategoryUsecase) {
	handler := categoryHandler{categoryUsecase}
	r.Get("/category", handler.GetCategories)
	r.Get("/category/:id_kategori", handler.GetCategoryByID)
	r.Post("/category", middleware.Auth, handler.AddCategory)
	r.Delete("/category/:id_kategori", middleware.Auth, handler.DeleteCategoryByID)
	r.Put("/category/:id_kategori", middleware.Auth, handler.UpdateCategory)
}

func (h *categoryHandler) AddCategory(c *fiber.Ctx) error {
	errs := []string{}
	userRole := c.Locals("user_role").(string)

	category := model.AddCategoryRequest{}
	err := c.BodyParser(&category)
	if err != nil {
		errs = append(errs, err.Error())
	}

	lastInsertedID, err := h.categoryUsecase.CreateNewCategory(userRole, category)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  errs,
		Data:    lastInsertedID,
	})
}

func (h *categoryHandler) GetCategories(c *fiber.Ctx) error {
	errs := []string{}

	categories, err := h.categoryUsecase.GetCategories()
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to GET data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to GET data",
		Errors:  errs,
		Data:    categories,
	})
}

func (h *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	errs := []string{}

	categoryID, err := strconv.Atoi(c.Params("id_kategori"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	category, err := h.categoryUsecase.GetCategoryByID(categoryID)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to GET data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to GET data",
		Errors:  errs,
		Data:    category,
	})
}

func (h *categoryHandler) DeleteCategoryByID(c *fiber.Ctx) error {
	errs := []string{}
	userRole := c.Locals("user_role").(string)

	categoryID, err := strconv.Atoi(c.Params("id_kategori"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = h.categoryUsecase.DeleteCategoryByID(userRole, categoryID)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Failed to DELETE data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to DELETE data",
		Errors:  errs,
		Data:    "delete data succeed",
	})
}

func (h *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	errs := []string{}
	userRole := c.Locals("user_role").(string)

	categoryID, err := strconv.Atoi(c.Params("id_kategori"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	category := model.AddCategoryRequest{}
	err = c.BodyParser(&category)
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = h.categoryUsecase.UpdateCategory(userRole, categoryID, category)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Failed to DELETE data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to UPDATE data",
		Errors:  errs,
		Data:    "update data succeed",
	})
}
