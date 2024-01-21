package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	// "github.com/golang-jwt/jwt/"
	"github.com/saharat-vithchataya/noname/services"
)

type accountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) accountHandler {
	return accountHandler{accountService: accountService}
}

func (h accountHandler) OpenNewAccount(c *fiber.Ctx) error {
	account := services.CreateNewAccountScheme{}
	err := c.BodyParser(&account)
	if err != nil {
		return err
	}
	response, err := h.accountService.OpenNewAccount(account)
	if err != nil {
		return err
	}
	return c.JSON(response)
}

func (h accountHandler) GetAccount(c *fiber.Ctx) error {
	// user := c.Locals("user").(*jwtt.Token)
	// claims := user.Claims.(jwtt.MapClaims)
	// accountID := int(claims["account_id"].(float64))
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	accountID := claims["account_id"].(float64)

	response, err := h.accountService.GetAccount(int(accountID))
	if err != nil {
		return err
	}
	return c.JSON(response)
}
