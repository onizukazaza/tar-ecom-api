package validation

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"github.com/onizukazaza/tar-ecom-api/entities"
)

func SellerIDGetting(ctx *fiber.Ctx) (string, error) {
	sellerID, ok := ctx.Locals("userID").(string)
	if !ok || sellerID == "" {
		return "", fiber.NewError(http.StatusUnauthorized, "Unauthorized access: sellerID is missing or invalid")
	}

	role, ok := ctx.Locals("role").(string)   
	if !ok || entities.Role(role) != entities.RoleSeller {
		return "", fiber.NewError(http.StatusForbidden, "Access restricted to sellers only")
	}

	return sellerID, nil   
}

func BuyerIDGetting(ctx *fiber.Ctx) (string, error) {
    buyerID, ok := ctx.Locals("userID").(string)
    if !ok || buyerID == "" {
        return "", fiber.NewError(http.StatusUnauthorized, "Unauthorized access: buyerID is missing or invalid")
    }

    role, ok := ctx.Locals("role").(string)
    if !ok || entities.Role(role) != entities.RoleBuyer {
        return "", fiber.NewError(http.StatusForbidden, "Access restricted to buyers only")
    }

    return buyerID, nil
}
