package main

import (
	"fiber-templ-learn/config"
	"fiber-templ-learn/internal/home"
	"fiber-templ-learn/internal/vacancy"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)
func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Static("/public", "./public")
	
	app.Get("/metrics", monitor.New(monitor.Config{
		Title: "Monitoring system",
		Refresh: time.Duration(1000),
	},
		))

	home.NewHandler(app)
	vacancy.NewHandler(app)
	log.Info(dbConf)

	app.Listen(":3000")
}
