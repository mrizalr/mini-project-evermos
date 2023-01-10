package httphandler

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
	r.Post("/auth/login", handler.LoginUser)
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
		return c.Status(http.StatusBadRequest).JSON(model.Response{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    "Register Succeed",
	})
}

func (h *userHandler) LoginUser(c *fiber.Ctx) error {
	userRequest := model.UserLoginRequest{}
	errs := []string{}

	if err := c.BodyParser(&userRequest); err != nil {
		errs = append(errs, err.Error())
	}

	user, err := h.userUsecase.Login(userRequest)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(http.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    user,
	})
}
