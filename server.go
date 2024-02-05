package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/rs/zerolog"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	logger := setupLogger(context.Background(), filepath.Join(cfg.LogDir, "server.log"))

	engine := handlebars.New("./views", ".hbs")
	server := fiber.New(fiber.Config{
		Views: engine,
	})

	server.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Fiber Template",
			"Message": "Your page content",
		})
	})
	server.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Fiber Template",
			"Message": "Your page content",
		}, "layouts/main")
	})

	addrStr := fmt.Sprintf(":%d", cfg.Port)
	err := server.Listen(addrStr)
	if err != nil {
		logger.Fatal().Err(err).Msg("starting app server")
	}
}

func setupLogger(ctx context.Context, logFilePath string) *zerolog.Logger {
	var outWriter = os.Stdout
	//logFilePath := filepath.Join(dir,"server.log")
	if logFilePath != "" && logFilePath != "stdout" {
		file, err := os.OpenFile(logFilePath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			log.Fatalln(err)
		}
		outWriter = file
	}
	cout := zerolog.ConsoleWriter{Out: outWriter, TimeFormat: time.RFC822}
	cout.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	// uncomment to remove timestamp from logs
	//out.FormatTimestamp = func(i interface{}) string {
	//	return ""
	//}
	baseLogger := zerolog.New(cout).With().Timestamp().Logger()
	logCtx := baseLogger.WithContext(ctx)
	l := zerolog.Ctx(logCtx)

	return l
}
