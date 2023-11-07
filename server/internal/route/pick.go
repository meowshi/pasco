package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	"github.com/meowshi/pasco-server/internal/usecase"
)

func SetupPickRouter(app *fiber.App, db *sqlx.DB) {
	storage := postgres.NewPickPostgresStorage(db)
	usecase := usecase.NewPickUsecase(storage)
	controller := controller.NewPickController(usecase)

	router := app.Group("/pick")

	router.Get("", controller.GetPickHistory)
	router.Post("", controller.CreatePick)
}
