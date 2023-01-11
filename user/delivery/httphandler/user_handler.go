package httphandler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/domain"
	"github.com/mrizalr/mini-project-evermos/middleware"
	"github.com/mrizalr/mini-project-evermos/model"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(r fiber.Router, userUsecase domain.UserUsecase) {
	handler := userHandler{userUsecase}
	r.Post("/auth/register", handler.RegisterUser)
	r.Post("/auth/login", handler.LoginUser)
	r.Get("/user", middleware.Auth, handler.GetMyProfile)
	r.Put("/user", middleware.Auth, handler.UpdateMyProfile)
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

func (h *userHandler) GetMyProfile(c *fiber.Ctx) error {
	errs := []string{}

	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	user, err := h.userUsecase.GetMyProfile(userID)
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
		Data:    user,
	})
}

func (h *userHandler) UpdateMyProfile(c *fiber.Ctx) error {
	user := model.UpdateUserRequest{}
	errs := []string{}

	userID, err := strconv.Atoi(c.Locals("user_id").(string))
	if err != nil {
		errs = append(errs, err.Error())
	}

	if err := c.BodyParser(&user); err != nil {
		errs = append(errs, err.Error())
	}

	err = h.userUsecase.UpdateMyProfile(userID, user)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadGateway).JSON(model.Response{
			Status:  false,
			Message: "Failed to Update data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Succeed to Update data",
		Errors:  nil,
		Data:    "Update user succeed",
	})
}
