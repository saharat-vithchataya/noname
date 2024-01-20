package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saharat-vithchataya/noname/services"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthenticationHandler(authService services.AuthService) authHandler {
	return authHandler{authService: authService}
}

func (h authHandler) Login(c *fiber.Ctx) error {

	userCredentials := services.LoginScheme{}
	err := c.BodyParser(&userCredentials)
	if err != nil {
		return err
	}

	response, err := h.authService.Login(userCredentials)
	if err != nil {
		return err
	}
	return c.JSON(response)
}
