package usecase

import "github.com/meowshi/pasco-server/pkg/domain"

type RegistrationUsecase interface {
	Update(registration *domain.Registration) error
}
