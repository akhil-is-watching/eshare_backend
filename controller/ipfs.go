package controller

import (
	"github.com/akhil-is-watching/encryptedFileSharing/service"
	"github.com/akhil-is-watching/encryptedFileSharing/types"
	"github.com/gofiber/fiber/v2"
)

type IPFSController struct {
	ipfsService *service.PinataIPFS
}

func NewIPFSController(ipfsService *service.PinataIPFS) *IPFSController {
	return &IPFSController{ipfsService: ipfsService}
}

func (c *IPFSController) PublishFile(ctx *fiber.Ctx) error {
	var req types.DocumentPackage
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ipfsHash, err := c.ipfsService.Publish(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"ipfsHash": ipfsHash})
}

func (c *IPFSController) RetrieveFile(ctx *fiber.Ctx) error {
	cid := ctx.Params("cid")
	if cid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "CID is required"})
	}

	document, err := c.ipfsService.Retrieve(cid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(document)
}
