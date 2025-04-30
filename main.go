package main

import (
	"fmt"
	"log"
	"os"

	"github.com/czlucius/traefik-dynamic-mux/dynamic"
	"github.com/czlucius/traefik-dynamic-mux/muxfile"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{
			AppName: "traefik-dynamic-mux",
		},
	)

	// Default Muxfile location is /etc/traefik/Muxfile
	// You can set the Muxfile location using the environment variable MUXFILE
	muxFile := os.Getenv("MUXFILE")
	if muxFile == "" {
		muxFile = "/etc/traefik/Muxfile"
	}
	fmt.Println("Muxfile:", muxFile)

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Traefik Configuration Manager!")
	})

	app.Get("/mux", func(c *fiber.Ctx) error {
		fmt.Println("Request received")
		config := &dynamic.Configuration{}
		// Exec the configuration from the Muxfile
		// Read the Muxfile with os
		b, err := os.ReadFile(muxFile)
		if err != nil {
			fmt.Print(err)
		}
		// Parse the Muxfile
		muxfile.ExecMuxFile(string(b), config)
		// Return the configuration as JSON
		return c.JSON(config)
	})

	log.Fatal(app.Listen(":9393"))
}
