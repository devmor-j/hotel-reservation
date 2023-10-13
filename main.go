package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", handleUser)

	app.Listen(*listenAddr)
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "James Bond"})
}
