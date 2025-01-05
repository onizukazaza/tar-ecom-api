package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onizukazaza/tar-ecom-api/pkg/custom"
	_userModel "github.com/onizukazaza/tar-ecom-api/pkg/user/model"
	_userService "github.com/onizukazaza/tar-ecom-api/pkg/user/service"
	"net/http"
)

type userControllerImpl struct {
	userService _userService.UserService
}

func NewUserControllerImpl(
	userService _userService.UserService,
) UserController {
	return &userControllerImpl{userService}
}

func (c *userControllerImpl) Listing(ctx *fiber.Ctx) error {
	userModelList, err := c.userService.Listing()
	if err != nil {
		// return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(userModelList)
}

func (c *userControllerImpl) CreateUser(ctx *fiber.Ctx) error {
    req := new(_userModel.CreateUserReq)
    customReq := custom.NewCustomFiberRequest(ctx)

    if err := customReq.Bind(req); err != nil {
        return custom.CustomError(ctx, http.StatusBadRequest, err.Error())
    }

    // ตรวจสอบว่าผู้ใช้งานมีอยู่ในระบบแล้วหรือไม่
    exists, err := c.userService.IsUserExists(req.Email)
    if err != nil {
        return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
    }
    if exists {
        return custom.CustomError(ctx, http.StatusConflict, "User already exists")
    }

    createdUser, err := c.userService.CreateUser(req)
    if err != nil {
        return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusCreated).JSON(createdUser)
}

func (c *userControllerImpl) FindUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := c.userService.FindUserByID(id)
	if err != nil {
		return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

func (c *userControllerImpl) EditUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := new(_userModel.EditUserReq)
	req.ID = id

	customReq := custom.NewCustomFiberRequest(ctx)
	if err := customReq.Bind(req); err != nil {
		return custom.CustomError(ctx, http.StatusBadRequest, err.Error())
	}

	err := c.userService.EditUser(req)
	if err != nil {
		return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}
