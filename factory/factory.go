package factory

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	categoryHandler "github.com/mrizalr/mini-project-evermos/category/delivery/httphandler"
	categoryRepository "github.com/mrizalr/mini-project-evermos/category/repository/mysql"
	categoryUsecase "github.com/mrizalr/mini-project-evermos/category/usecase"
	productHandler "github.com/mrizalr/mini-project-evermos/product/delivery/httphandler"
	productRepository "github.com/mrizalr/mini-project-evermos/product/repository/mysql"
	productUsecase "github.com/mrizalr/mini-project-evermos/product/usecase"
	provinceHandler "github.com/mrizalr/mini-project-evermos/province/delivery/httphandler"
	storeHandler "github.com/mrizalr/mini-project-evermos/store/delivery/httphandler"
	storeRepository "github.com/mrizalr/mini-project-evermos/store/repository/mysql"
	storeUsecase "github.com/mrizalr/mini-project-evermos/store/usecase"
	userHandler "github.com/mrizalr/mini-project-evermos/user/delivery/httphandler"
	userRepository "github.com/mrizalr/mini-project-evermos/user/repository/mysql"
	userUsecase "github.com/mrizalr/mini-project-evermos/user/usecase"
)

func Init(r fiber.Router, db *gorm.DB) {
	v1 := r.Group("/api/v1")

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

	mysqlProductRepository := productRepository.NewMysqlProductRepository(db)
	productUsecase := productUsecase.NewProductUsecase(mysqlProductRepository)
	productHandler.NewProductHandler(v1, productUsecase)
}
