package httphandler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/middleware"
	"github.com/mrizalr/mini-project-evermos/model"
)

type storeHandler struct {
	storeUsecase domain.StoreUsecase
}

func NewStoreHandler(r fiber.Router, storeUsecase domain.StoreUsecase) {
	handler := storeHandler{storeUsecase}
	r.Get("/toko", handler.GetStores)
	r.Get("/toko/my", middleware.Auth, handler.GetMyStore)
	r.Get("/toko/:id_toko", handler.GetStoreByID)
	r.Put("/toko/:id_toko", middleware.Auth, handler.UpdateMyStore)
}

func (h *storeHandler) GetMyStore(c *fiber.Ctx) error {
	errs := []string{}

	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	store, err := h.storeUsecase.GetMyStore(userID)
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
		Errors:  nil,
		Data:    store,
	})
}

func (h *storeHandler) UpdateMyStore(c *fiber.Ctx) error {
	errs := []string{}
	request := model.UpdateStoreRequest{}
	storagePath := "images/store_photo/"

	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	storeIDparam := c.Params("id_toko")
	storeID, err := strconv.Atoi(storeIDparam)
	if err != nil {
		errs = append(errs, err.Error())
	}

	store, err := h.storeUsecase.GetStoreByID(storeID)
	if err != nil {
		errs = append(errs, err.Error())
	}

	request.Name = c.FormValue("nama_toko")
	file, _ := c.FormFile("photo")

	if file != nil {
		if store.PhotoURL != "" {
			os.Remove(storagePath + store.PhotoURL)
		}

		request.PhotoURL = fmt.Sprintf("%d_%s", storeID, file.Filename)

		relativePath := storagePath + request.PhotoURL
		err = c.SaveFile(file, relativePath)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	err = h.storeUsecase.UpdateStore(userID, storeID, request)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to UPDATE data",
		Errors:  nil,
		Data:    "Update toko succeed",
	})
}

func (h *storeHandler) GetStoreByID(c *fiber.Ctx) error {
	errs := []string{}
	storeIDparam := c.Params("id_toko")
	storeID, err := strconv.Atoi(storeIDparam)
	if err != nil {
		errs = append(errs, err.Error())
	}

	store, err := h.storeUsecase.GetStoreByID(storeID)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to UPDATE data",
		Errors:  nil,
		Data:    store,
	})
}

func (h *storeHandler) GetStores(c *fiber.Ctx) error {
	errs := []string{}

	opts := model.GetStoresOptions{}
	err := c.QueryParser(&opts)
	if err != nil {
		errs = append(errs, err.Error())
	}

	stores, err := h.storeUsecase.GetStores(opts)
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
		Errors:  nil,
		Data:    stores,
	})
}
