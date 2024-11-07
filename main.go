package main

import (
	"github.com/akhil-is-watching/encryptedFileSharing/controller"
	"github.com/akhil-is-watching/encryptedFileSharing/routes"
	"github.com/akhil-is-watching/encryptedFileSharing/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupCommonRoutes(app)
	routes.SetupIPFSRoutes(app, controller.NewIPFSController(service.NewPinataIPFS("cfd1a0d5f033336cd3a5", "dc79e6c518ad51e3b0f99ee2f9ae581a731a11ddc569067d3db2c2053d9d8d44")))

	app.Listen(":5000")
}
