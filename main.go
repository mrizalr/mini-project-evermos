package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mrizalr/mini-project-evermos/config"
	"github.com/mrizalr/mini-project-evermos/database"
	"github.com/mrizalr/mini-project-evermos/factory"
	"github.com/spf13/viper"
)

func main() {
	db := database.New()

	app := fiber.New()
	factory.Init(app, db)

	app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}
