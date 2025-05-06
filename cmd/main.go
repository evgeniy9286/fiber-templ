package main

import (
	"fiber-templ-learn/config"
	"fiber-templ-learn/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)
func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	
	app.Get("/metrics", monitor.New())

	home.NewHandler(app)
	log.Info(dbConf)

	app.Listen(":3000")
}
