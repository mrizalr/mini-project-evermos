package factory

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	addressHandler "github.com/mrizalr/mini-project-evermos/address/delivery/httphandler"
	addressRepository "github.com/mrizalr/mini-project-evermos/address/repository/mysql"
	addressUsecase "github.com/mrizalr/mini-project-evermos/address/usecase"
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
	transactionHandler "github.com/mrizalr/mini-project-evermos/transaction/delivery/httphandler"
	transactionRepository "github.com/mrizalr/mini-project-evermos/transaction/repository/mysql"
	transactionUsecase "github.com/mrizalr/mini-project-evermos/transaction/usecase"
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
	productUsecase := productUsecase.NewProductUsecase(mysqlProductRepository, mysqlStoreRepository, mysqlCategoryRepository)
	productHandler.NewProductHandler(v1, productUsecase, storeUsecase)

	mysqlAddressRepository := addressRepository.NewMysqlAddressRepository(db)
	addressUsecase := addressUsecase.NewAddressUsecase(mysqlAddressRepository)
	addressHandler.NewAddressHandler(v1, addressUsecase)

	transactionRepository := transactionRepository.NewTransactionRepository(db)
	transactionUsecase := transactionUsecase.NewTransactionUsecase(transactionRepository, mysqlProductRepository)
	transactionHandler.NewTransactionHandler(v1, transactionUsecase)
}
