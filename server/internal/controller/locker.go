package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/usecase"
	"github.com/sirupsen/logrus"
)

type LockerController struct {
	lockerUsecase usecase.LockerUsecase
	pickUsecase   usecase.PickUsecase
}

func NewLockerController(lu usecase.LockerUsecase, pu usecase.PickUsecase) *LockerController {
	return &LockerController{
		lockerUsecase: lu,
		pickUsecase:   pu,
	}
}

func (c *LockerController) GetLockerEvents(ctx *fiber.Ctx) error {
	events, err := c.lockerUsecase.GetLockerEvents()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(&events)
}

func (c *LockerController) Print(ctx *fiber.Ctx) error {
	pickId := ctx.QueryInt("pickId")

	req := &domain.PrintBraceletReq{}
	err := ctx.BodyParser(req)
	if err != nil {
		logrus.Errorf("Ошибка при парсинге тела PrintBracelets: %s.", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.lockerUsecase.PrintBracelet(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if pickId != 0 {
		c.pickUsecase.UpdateFromBracelet(int64(pickId), true)
	}

	return nil
}

func (c *LockerController) GetPrinters(ctx *fiber.Ctx) error {
	printers, err := c.lockerUsecase.GetPrinters()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(&printers)
}
