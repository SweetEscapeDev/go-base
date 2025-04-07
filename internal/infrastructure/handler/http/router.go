package http

import (
	"go-base/config"
	"go-base/internal/infrastructure/handler/http/routes"
	"go-base/internal/infrastructure/middleware"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(c *config.Config) {

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

	routes.HealthRouteInit(router)

	router.Run(c.HTTPServerAddress)
}
