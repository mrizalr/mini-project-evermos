package httphandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/middleware"
	"github.com/mrizalr/mini-project-evermos/model"
)

type transactionHandler struct {
	transactionUsecase domain.TransactionUsecase
}

func NewTransactionHandler(r fiber.Router, transactionUsecase domain.TransactionUsecase) {
	handler := transactionHandler{transactionUsecase}
	r.Post("/trx", middleware.Auth, handler.CreateNewTransaction)
}

func (h *transactionHandler) CreateNewTransaction(c *fiber.Ctx) error {
	errs := []string{}
	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	trxRequest := model.CreateTransactionRequest{}
	err = c.BodyParser(&trxRequest)
	if err != nil {
		errs = append(errs, err.Error())
	}

	lastInsertID, err := h.transactionUsecase.CreateNewTransaction(userID, trxRequest)
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
