package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/usecase"
)

type RegistrationController struct {
	registrationUsecase usecase.RegistrationUsecase
	pickUsecase         usecase.PickUsecase
}

func NewRegistrationController(ru usecase.RegistrationUsecase, pu usecase.PickUsecase) *RegistrationController {
	return &RegistrationController{
		registrationUsecase: ru,
		pickUsecase:         pu,
	}
}

func (c *RegistrationController) Update(ctx *fiber.Ctx) error {
	pickId := ctx.QueryInt("pickId")
	reg := &domain.Registration{}
	err := ctx.BodyParser(reg)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.registrationUsecase.Update(reg)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if pickId != 0 {
		err = c.pickUsecase.UpdateFromList(int64(pickId), reg.EventUuid, reg.Status, true)
	}

	return nil
}
