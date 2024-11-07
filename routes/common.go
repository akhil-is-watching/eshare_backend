package routes

import "github.com/gofiber/fiber/v2"

func SetupCommonRoutes(app *fiber.App) {
	common := app.Group("/common")
	common.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
	})
}
