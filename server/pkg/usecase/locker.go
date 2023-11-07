package usecase

import "github.com/meowshi/pasco-server/pkg/domain"

type LockerUsecase interface {
	GetLockerEvents() ([]*domain.LockerEvent, error)
	PrintBracelet(req *domain.PrintBraceletReq) error
	GetPrinters() ([]*domain.Printer, error)
}
