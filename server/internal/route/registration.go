package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	"github.com/meowshi/pasco-server/internal/usecase"
	"google.golang.org/api/sheets/v4"
)

func SetupRegistrationRouter(app *fiber.App, db *sqlx.DB, service *sheets.Service) {
	rs := postgres.NewRegistrationPostgresStorage(db)
	es := postgres.NewEnvPostgresStorage(db)
	ps := postgres.NewPickPostgresStorage(db)

	ru := usecase.NewRegistrationUsecase(rs, es, service)
	pu := usecase.NewPickUsecase(ps)

	c := controller.NewRegistrationController(ru, pu)

	router := app.Group("/registration")

	router.Patch("", c.Update)
}
