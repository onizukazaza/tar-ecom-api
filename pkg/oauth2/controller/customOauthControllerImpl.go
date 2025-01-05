package controller

import (
	"github.com/gofiber/fiber/v2"
	_oauth2Service "github.com/onizukazaza/tar-ecom-api/pkg/oauth2/service"
	"github.com/onizukazaza/tar-ecom-api/pkg/oauth2/model"
)


type oauth2ControllerImpl struct {
	oauth2Service _oauth2Service.OAuth2Service
	secretKey     string

}

func NewOAuth2Controller(oauth2Service _oauth2Service.OAuth2Service, secretKey string) OAuth2Controller {
	return &oauth2ControllerImpl{
		oauth2Service: oauth2Service,
	    secretKey:     secretKey,

	}
}

func (c *oauth2ControllerImpl) Login(ctx *fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, err := c.oauth2Service.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"token": token})
}

func (c *oauth2ControllerImpl) Logout(ctx *fiber.Ctx) error {
    tokenString, err := ExtractTokenFromHeader(ctx)
    if err != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
    }

    err = c.oauth2Service.Logout(tokenString)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(fiber.Map{"message": "logout successful"})
}
