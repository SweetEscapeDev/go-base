package initializer

import (
	"go-base/config"

	"go-base/internal/infrastructure/middleware"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func InitializeServer(cfg config.Config) {
	InitializeMidtrans(cfg)
	db := InitializeDatabase(cfg)
	// InitializeGORMLog(db)

	validate := validator.New()

	router := gin.New()
	router.Use(gin.Recovery())

	cors := config.NewCors()
	router.Use(cors)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context:    ginzap.Fn(middleware.GinzapExtraFields),
	}))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	InitializeHealth(router)
	InitializeExample(router, db, validate)

	router.Run(cfg.HTTPServerAddress)
}
