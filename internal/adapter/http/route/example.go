package route

import (
	"go-base/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleRoute interface {
	RegisterRoutes()
	Example(c *gin.Context)
}

type exampleRouteImpl struct {
	router         *gin.Engine
	exampleUsecase usecase.ExampleUseCase
}

func ProviderExamplehRoute(router *gin.Engine, exampleUsecase usecase.ExampleUseCase) ExampleRoute {
	route := &exampleRouteImpl{
		router,
		exampleUsecase,
	}
	route.RegisterRoutes()
	return route
}

func (r *exampleRouteImpl) RegisterRoutes() {
	r.router.GET("/example", r.Example)
}

func (r *exampleRouteImpl) Example(c *gin.Context) {
	result, _ := r.exampleUsecase.ExampleCase(c)
	c.JSON(http.StatusOK, gin.H{
		"status":  result,
		"message": "Service is running",
	})
}
