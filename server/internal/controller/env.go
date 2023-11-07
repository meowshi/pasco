package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/usecase"
)

type EnvController struct {
	usecase usecase.EnvUsecase
}

func NewEnvController(u usecase.EnvUsecase) *EnvController {
	return &EnvController{
		usecase: u,
	}
}

func (c *EnvController) Get(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
	if len(key) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Key param dosen't provided.")
	}

	env, err := c.usecase.Get(key)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(env)
}

func (c *EnvController) Update(ctx *fiber.Ctx) error {
	env := &domain.Env{}
	err := ctx.BodyParser(env)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.usecase.Update(env)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *EnvController) UpdateSheetTitle(ctx *fiber.Ctx) error {
	err := c.usecase.UpdateSheetTitle()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
