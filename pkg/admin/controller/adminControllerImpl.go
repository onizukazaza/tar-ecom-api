package controller

import (
	"github.com/gofiber/fiber/v2"
	_adminService "github.com/onizukazaza/tar-ecom-api/pkg/admin/service"
)

type adminControllerImpl struct {
	adminService _adminService.AdminService
}

func NewAdminControllerImpl(
   adminService _adminService.AdminService,
) AdminController {
	return &adminControllerImpl{adminService}
}

func (c *adminControllerImpl) Listing(ctx *fiber.Ctx) error {
	userModelList , err := c.adminService.Listing()
	if err!= nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }
	return ctx.Status(fiber.StatusOK).JSON(userModelList)
}
