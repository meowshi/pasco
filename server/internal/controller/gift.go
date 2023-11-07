package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meowshi/pasco-server/pkg/usecase"
)

type GiftController struct {
	giftUsecase usecase.GiftUsecase
	pickUsecase usecase.PickUsecase
}

func NewGiftController(gu usecase.GiftUsecase, pu usecase.PickUsecase) *GiftController {
	return &GiftController{
		giftUsecase: gu,
		pickUsecase: pu,
	}
}

func (c *GiftController) Get(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
	if len(key) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "'key' is not set")
	}

	res, err := c.giftUsecase.Get(key)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(res)
}
func (c *GiftController) Give(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
	count := ctx.Query("count")
	pickId := ctx.QueryInt("pickId")

	login := ctx.Params("login")

	if len(key) == 0 || len(count) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Provide both 'key' and 'count' query params.")
	}
	if len(login) == 0 {
		return fiber.NewError(fiber.StatusNotFound)
	}

	err := c.giftUsecase.Give(key, login, count)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if pickId != 0 {
		c.pickUsecase.UpdateFromGift(int64(pickId), true)
	}

	return nil
}
