package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
    "errors"
)

// JWT Middleware for token validation
func JWTMiddleware(secretKey string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString, err := ExtractTokenFromHeader(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			// ตรวจสอบว่าข้อผิดพลาดเกี่ยวกับ Token หมดอายุหรือไม่
			if errors.Is(err, jwt.ErrTokenExpired) {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token expired"})
			}
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		if !token.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		// บันทึก Claims ลงใน Context
		ctx.Locals("userID", claims["id"])
		ctx.Locals("role", claims["role"])
		return ctx.Next()
	}
}

// Extract Token from Authorization Header
func ExtractTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("missing token")
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		return tokenString[7:], nil
	}
	return "", fmt.Errorf("invalid token format")
}


// Role Authorizing Middleware
func RoleAuthorizing(ctx *fiber.Ctx, requiredRole string) error {
	role := ctx.Locals("role")
	if role != requiredRole {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": fmt.Sprintf("access restricted to %s", requiredRole)})
	}
	return ctx.Next()
}


func (c *oauth2ControllerImpl) UserAuthorizing(ctx *fiber.Ctx) error {
	tokenString, err := ExtractTokenFromHeader(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
	}

	// Validate the token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.secretKey), nil
	})
	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired token"})
	}

	// Save claims to context
	ctx.Locals("userID", claims["id"])
	ctx.Locals("role", claims["role"])
	return ctx.Next()
}
