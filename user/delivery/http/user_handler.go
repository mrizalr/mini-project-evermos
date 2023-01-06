package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/model"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(r fiber.Router, userUsecase domain.UserUsecase) {
	handler := userHandler{userUsecase}
	r.Post("/auth/register", handler.RegisterUser)
}

func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	user := model.UserRegisterRequest{}
	errs := []string{}

	if err := c.BodyParser(&user); err != nil {
		errs = append(errs, err.Error())
	}

	if err := h.userUsecase.Register(user); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		c.Status(http.StatusBadRequest)
		return c.JSON(model.Response{
			Status:  "false",
			Message: "Failed to POST data",
			Errors:  errs,
			Data:    nil,
		})
	}

	c.Status(http.StatusCreated)
	return c.JSON(model.Response{
		Status:  "true",
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    "Register Succeed",
	})
}
