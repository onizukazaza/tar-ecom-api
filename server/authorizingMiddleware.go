package server

import (
	"github.com/gofiber/fiber/v2"
	_oauth2Controller "github.com/onizukazaza/tar-ecom-api/pkg/oauth2/controller"
)

type authorizingMiddleware struct {
    _oauth2Controller.OAuth2Controller
}

func (m *authorizingMiddleware) MiddlewareFunc() fiber.Handler {
    return func(ctx *fiber.Ctx) error {

        return m.OAuth2Controller.UserAuthorizing(ctx)
    }
}


