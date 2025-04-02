package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRouteInterface interface {
	RegisterRoutes()
	HealthCheck(c *gin.Context)
}

type HealthRoute struct {
	router *gin.Engine
}

func HealthRouteInit(router *gin.Engine) HealthRouteInterface {
	route := &HealthRoute{
		router: router,
	}
	route.RegisterRoutes()
	return route
}

func (r *HealthRoute) RegisterRoutes() {
	r.router.GET("/health", r.HealthCheck)
}

func (r *HealthRoute) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"message": "Service is running",
	})
}
