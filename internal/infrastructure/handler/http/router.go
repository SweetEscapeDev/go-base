package http

import (
	"go-base/config"
	"go-base/internal/infrastructure/handler/http/routes"
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
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	routes.HealthRouteInit(router)
	routes.WebhookbRouteInit(router)

	router.Run(c.HTTPServerAddress)
}
