package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	categoryHandler "github.com/mrizalr/mini-project-evermos/category/delivery/httphandler"
	categoryRepository "github.com/mrizalr/mini-project-evermos/category/repository/mysql"
	categoryUsecase "github.com/mrizalr/mini-project-evermos/category/usecase"
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
	db.AutoMigrate(&domain.User{}, &domain.Store{}, &domain.Category{})

	app := fiber.New()
	v1 := app.Group("/api/v1")

	mysqlUserRepository := userRepository.NewMysqlUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(mysqlUserRepository)
	userHandler.NewUserHandler(v1, userUsecase)

	provinceHandler.NewProvinceHandler(v1)

	mysqlStoreRepository := storeRepository.NewMysqlStoreRepository(db)
	storeUsecase := storeUsecase.NewStoreUsecase(mysqlStoreRepository)
	storeHandler.NewStoreHandler(v1, storeUsecase)

	mysqlCategoryRepository := categoryRepository.NewMysqlCategoryRepository(db)
	categoryUsecase := categoryUsecase.NewCategoryUsecase(mysqlCategoryRepository)
	categoryHandler.NewCategoryHandler(v1, categoryUsecase)

	app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}
