package main

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

type config struct {
	Port   int    `env:"PORT" envDefault:"4000"`
	LogDir string `env:"LOGDIR,expand" envDefault:"${HOME}/tmp"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("problem parsing config: %+v", err)
	}

	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It is alive!")
	})

	addrStr := fmt.Sprintf(":%d", cfg.Port)
	server.Listen(addrStr)
}
