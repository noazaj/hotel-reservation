package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zajicekn/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Noah",
		LastName:  "Dingle",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Brian")
}
