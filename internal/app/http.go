package app

import (
	"gorm.io/gorm"

	"github.com/bangadam/backend-test-majoo/cmd/docs"
	"github.com/bangadam/backend-test-majoo/internal/adapter/inbound/auth_handler"
	"github.com/bangadam/backend-test-majoo/internal/adapter/inbound/report_handler"
	"github.com/bangadam/backend-test-majoo/internal/adapter/outbound/merchant_repository"
	"github.com/bangadam/backend-test-majoo/internal/adapter/outbound/outlet_repository"
	"github.com/bangadam/backend-test-majoo/internal/adapter/outbound/transaction_repository"
	"github.com/bangadam/backend-test-majoo/internal/adapter/outbound/user_repository"
	"github.com/bangadam/backend-test-majoo/internal/service/auth_service"
	"github.com/bangadam/backend-test-majoo/internal/service/report_transaction_service"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
)

type Handlers struct {
	Mysql      *gorm.DB
	R             *fiber.App
	Logger        *zap.Logger
}

func (h *Handlers) SetupRouter() {
	h.R.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})
	h.R.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Talk Is Cheap Show Me Your Code")
	})

	// Swagger Documection
	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	docs.SwaggerInfo.Host =  serverHost + ":" + serverPort 
	h.R.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	//initialize Repository
	// productCategoryRep := productcategoryrps.NewProductCategoryPostgres(h.Postgres)
	userRepository := user_repository.NewUserMysql(h.Mysql)
	transactionRepository := transaction_repository.NewTransactionRepository(h.Mysql)
	merchantRepository := merchant_repository.NewMerchantRepository(h.Mysql)
	outletRepository := outlet_repository.NewOutletRepository(h.Mysql)

	//initialize bussiness
	// productCategoryService := productcategorysvc.NewProductCategoryService(h.Logger, productCategoryRep)
	authService := auth_service.NewAuthService(h.Logger, userRepository)
	reportTransactionService := report_transaction_service.NewReportTransactionService(h.Logger, merchantRepository,outletRepository, transactionRepository)

	//handlers initialize
	// productcategoryhdl.NewProductCategoryHandler(h.R, productCategoryService)
	auth_handler.NewAuthHandler(h.R, authService)
	report_handler.NewReportHandler(h.R, reportTransactionService)
}
