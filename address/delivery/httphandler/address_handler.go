package httphandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/middleware"
	"github.com/mrizalr/mini-project-evermos/model"
)

type addressHandler struct {
	addressUsecase domain.AddressUsecase
}

func NewAddressHandler(r fiber.Router, addressUsecase domain.AddressUsecase) {
	handler := addressHandler{addressUsecase}
	r.Post("/user/alamat", middleware.Auth, handler.CreateNewAddress)
	r.Get("/user/alamat", middleware.Auth, handler.GetMyAddress)
	r.Get("/user/alamat/:id", middleware.Auth, handler.GetAddressByID)
	r.Put("/user/alamat/:id", middleware.Auth, handler.UpdateAddress)
	r.Delete("/user/alamat/:id", middleware.Auth, handler.DeleteAddress)
}

func (h *addressHandler) CreateNewAddress(c *fiber.Ctx) error {
	errs := []string{}
	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	address := domain.Address{}
	err = c.BodyParser(&address)
	if err != nil {
		errs = append(errs, err.Error())
	}

	address.UserID = uint(userID)
	lastInsertID, err := h.addressUsecase.CreateNewAddress(address)
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
		Status:  true,
		Message: "Success to POST data",
		Errors:  errs,
		Data:    lastInsertID,
	})
}

func (h *addressHandler) GetMyAddress(c *fiber.Ctx) error {
	errs := []string{}
	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	address, err := h.addressUsecase.GetMyAddress(userID)
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
		Message: "Success to GET data",
		Errors:  errs,
		Data:    address,
	})
}

func (h *addressHandler) GetAddressByID(c *fiber.Ctx) error {
	errs := []string{}
	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	address, err := h.addressUsecase.GetAddressByID(userID, addressID)
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
		Message: "Success to GET data",
		Errors:  errs,
		Data:    address,
	})
}

func (h *addressHandler) UpdateAddress(c *fiber.Ctx) error {
	errs := []string{}
	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	address := model.UpdateAddressRequest{}
	err = c.BodyParser(&address)
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = h.addressUsecase.UpdateAddress(userID, addressID, address)
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
		Message: "Success to UPDATE data",
		Errors:  errs,
		Data:    "",
	})
}

func (h *addressHandler) DeleteAddress(c *fiber.Ctx) error {
	errs := []string{}
	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	addressID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errs = append(errs, err.Error())
	}

	err = h.addressUsecase.DeleteAddress(userID, addressID)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to DELETE data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success to DELETE data",
		Errors:  errs,
		Data:    "",
	})
}
