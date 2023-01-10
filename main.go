package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mrizalr/mini-project-evermos/config"
	"github.com/mrizalr/mini-project-evermos/database"
	"github.com/mrizalr/mini-project-evermos/domain"
	provinceHandler "github.com/mrizalr/mini-project-evermos/province/delivery/httphandler"
	storeHandler "github.com/mrizalr/mini-project-evermos/store/delivery/httphandler"
	storeRepository "github.com/mrizalr/mini-project-evermos/store/repository/mysql"
	storeUsecase "github.com/mrizalr/mini-project-evermos/store/usecase"
	userHandler "github.com/mrizalr/mini-project-evermos/user/delivery/httphandler"
	userRepository "github.com/mrizalr/mini-project-evermos/user/repository/mysql"
	userUsecase "github.com/mrizalr/mini-project-evermos/user/usecase"
	"github.com/spf13/viper"
)

func main() {
	db, err := database.New()
	if err != nil {
		panic(fmt.Errorf("cannot connect to database : %v", err.Error()))
	}

	app := fiber.New()
	v1 := app.Group("/api/v1")

	mysqlUserRepository := userRepository.NewMysqlUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(mysqlUserRepository)
	userHandler.NewUserHandler(v1, userUsecase)

	provinceHandler.NewProvinceHandler(v1)

	mysqlStoreRepository := storeRepository.NewMysqlStoreRepository(db)
	storeUsecase := storeUsecase.NewStoreUsecase(mysqlStoreRepository)
	storeHandler.NewStoreHandler(v1, storeUsecase)

	db.AutoMigrate(&domain.User{}, &domain.Store{})
	app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}
