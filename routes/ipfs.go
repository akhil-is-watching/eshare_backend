package routes

import (
	"github.com/akhil-is-watching/encryptedFileSharing/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupIPFSRoutes(app *fiber.App, controller *controller.IPFSController) {
	ipfs := app.Group("/ipfs")
	ipfs.Post("/publish", controller.PublishFile)
	ipfs.Get("/retrieve/:cid", controller.RetrieveFile)
}
