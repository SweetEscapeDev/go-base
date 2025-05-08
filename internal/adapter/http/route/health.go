package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRoute interface {
	RegisterRoutes()
	HealthCheck(c *gin.Context)
}

type healthRouteImpl struct {
	router *gin.Engine
}

func ProviderHealthRoute(router *gin.Engine) HealthRoute {
	route := &healthRouteImpl{
		router: router,
	}
	route.RegisterRoutes()
	return route
}

func (r *healthRouteImpl) RegisterRoutes() {
	r.router.GET("/health", r.HealthCheck)
}

func (r *healthRouteImpl) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"message": "Service is running",
	})
}
