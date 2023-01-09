package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mrizalr/mini-project-evermos/config"
	"github.com/mrizalr/mini-project-evermos/database"
	"github.com/mrizalr/mini-project-evermos/domain"
	provinceHandler "github.com/mrizalr/mini-project-evermos/province/delivery/http"
	userHandler "github.com/mrizalr/mini-project-evermos/user/delivery/http"
	"github.com/mrizalr/mini-project-evermos/user/repository/mysql"
	"github.com/mrizalr/mini-project-evermos/user/usecase"
	"github.com/spf13/viper"
)

func main() {
	db, err := database.New()
	if err != nil {
		panic(fmt.Errorf("cannot connect to database : %v", err.Error()))
	}

	app := fiber.New()
	v1 := app.Group("/api/v1")

	mysqlUserRepository := mysql.NewMysqlUserRepository(db)
	userUsecase := usecase.NewUserUsecase(mysqlUserRepository)
	userHandler.NewUserHandler(v1, userUsecase)

	provinceHandler.NewProvinceHandler(v1)

	db.AutoMigrate(&domain.User{})
	app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}
