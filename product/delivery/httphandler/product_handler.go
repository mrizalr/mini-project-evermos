package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
)

type productHandler struct {
	productUsecase domain.ProductUsecase
}

func NewProductHandler(r fiber.Router, productUsecase domain.ProductUsecase) {
	handler := productHandler{productUsecase}
	r.Post("/product", handler.CreateNewProduct)
}

func (h *productHandler) CreateNewProduct(c *fiber.Ctx) error {
	return nil
}
