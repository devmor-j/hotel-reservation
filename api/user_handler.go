package api

import (
	"github.com/devmor-j/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("James")
}

func HandleGetUser(c *fiber.Ctx) error {
	user := types.User{
		Firstname: "me",
		Lastname:  "gmail",
	}
	return c.JSON(user)
}
