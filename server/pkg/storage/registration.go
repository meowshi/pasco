package storage

import (
	"github.com/meowshi/pasco-server/pkg/domain"
)

type RegistrationStorage interface {
	Create(registration *domain.Registration) error
	CreateMultiple(registrations ...*domain.Registration) error
	Update(registration *domain.Registration) error
}
