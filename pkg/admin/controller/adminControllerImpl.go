package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	_adminModel "github.com/onizukazaza/tar-ecom-api/pkg/admin/model"
	_adminService "github.com/onizukazaza/tar-ecom-api/pkg/admin/service"
	"github.com/onizukazaza/tar-ecom-api/pkg/validation"
	"github.com/onizukazaza/tar-ecom-api/pkg/custom"
)

type adminControllerImpl struct {
	adminService _adminService.AdminService
}

func NewAdminControllerImpl(adminService _adminService.AdminService) AdminController {
	return &adminControllerImpl{adminService}
}

func (c *adminControllerImpl) SetRole(ctx *fiber.Ctx) error {
	// ตรวจสอบว่าเป็น Admin หรือไม่
	adminID, err := validation.AdminIDGetting(ctx)
	if err != nil {
		return custom.CustomError(ctx, http.StatusUnauthorized, err.Error())
	}

	// ดึง ID จาก URL Param
	userID := ctx.Params("id")
	if userID == "" {
		return custom.CustomError(ctx, http.StatusBadRequest, "Missing user ID in URL")
	}

	// ดึงข้อมูล Role จาก Request Body
	var req _adminModel.SetRoleReq
	if err := ctx.BodyParser(&req); err != nil {
		return custom.CustomError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	// ผสม ID จาก URL Param กับข้อมูลใน Body
	req.ID = userID

	// เรียก Service เพื่อเปลี่ยน Role
	err = c.adminService.SetRole(&req)
	if err != nil {
		return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User role updated successfully",
		"admin_id": adminID,
	})
}

