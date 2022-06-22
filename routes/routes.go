package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leogouveia/go-fiber-mongo/controllers"
)

func HandleRequests() {
	app := fiber.New()

	app.Get("/", controllers.Index)
	app.Get("/realize/:id", controllers.GetRealize)
	app.Get("/realize", controllers.GetAllRealize)

	app.Listen(":3001")
}
