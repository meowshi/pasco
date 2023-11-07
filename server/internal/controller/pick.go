package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/usecase"
	"github.com/sirupsen/logrus"
)

type PickController struct {
	usecase usecase.PickUsecase
}

func NewPickController(usecase usecase.PickUsecase) *PickController {
	return &PickController{
		usecase: usecase,
	}
}

func (c *PickController) GetPickHistory(ctx *fiber.Ctx) error {
	picks, err := c.usecase.GetPickHistory()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(picks)
}

func (c *PickController) CreatePick(ctx *fiber.Ctx) error {
	req := &domain.CreatePickReq{}
	err := ctx.BodyParser(req)
	if err != nil {
		logrus.Errorf("Ошибка при unmarshall тела CreatePickReq: %s.", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := c.usecase.CreatePick(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendString(fmt.Sprintf("%d", id))
}
