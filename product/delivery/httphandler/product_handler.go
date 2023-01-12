package httphandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/middleware"
	"github.com/mrizalr/mini-project-evermos/model"
)

type productHandler struct {
	productUsecase domain.ProductUsecase
	storeUsecase   domain.StoreUsecase
}

func NewProductHandler(r fiber.Router, productUsecase domain.ProductUsecase, storeUsecase domain.StoreUsecase) {
	handler := productHandler{productUsecase, storeUsecase}
	r.Post("/product", middleware.Auth, handler.CreateNewProduct)
	r.Get("/product/", handler.GetProducts)
	r.Get("/product/:id", handler.GetProductByID)
}

func (h *productHandler) CreateNewProduct(c *fiber.Ctx) error {
	errs := []string{}
	basePathURL := "images/product_photo/"

	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	store, err := h.storeUsecase.GetMyStore(userID)
	if err != nil {
		errs = append(errs, err.Error())
	}

	productResellerPrice, err := strconv.ParseFloat(c.FormValue("harga_reseller"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}

	productConsumentPrice, err := strconv.ParseFloat(c.FormValue("harga_konsumen"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}

	productStock, err := strconv.Atoi(c.FormValue("stok"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	productCategoryID, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	createProductRequest := model.CreateProductRequest{
		Name:           c.FormValue("nama_produk"),
		ResellerPrice:  float32(productResellerPrice),
		ConsumentPrice: float32(productConsumentPrice),
		Stock:          productStock,
		Description:    c.FormValue("deskripsi"),
		CategoryID:     uint(productCategoryID),
		Photos:         []string{},
	}

	photosFile, _ := c.MultipartForm()
	for _, fileHeader := range photosFile.File {
		for _, header := range fileHeader {
			relativePath := fmt.Sprintf("%d_%s", store.ID, header.Filename)
			err := c.SaveFile(header, basePathURL+relativePath)
			if err != nil {
				errs = append(errs, err.Error())
			}
			createProductRequest.Photos = append(createProductRequest.Photos, relativePath)
		}
	}

	lastInsertID, err := h.productUsecase.CreateNewProduct(int(store.ID), createProductRequest)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  false,
		Message: "Succeed to POST data",
		Errors:  errs,
		Data:    lastInsertID,
	})
}

func (h *productHandler) GetProductByID(c *fiber.Ctx) error {
	errs := []string{}
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	product, err := h.productUsecase.GetProductByID(productID)
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
		Status:  false,
		Message: "Succeed to GET data",
		Errors:  errs,
		Data:    product,
	})
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	errs := []string{}

	product, err := h.productUsecase.GetProducts()
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
		Status:  false,
		Message: "Succeed to GET data",
		Errors:  errs,
		Data:    product,
	})
}
