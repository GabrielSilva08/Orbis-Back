package tagsController

import "github.com/gofiber/fiber/v2"

type TagControllerInterface interface {
	Create(ctx *fiber.Ctx) error
	ListAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
