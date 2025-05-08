//go:build wireinject
// +build wireinject

package initializer

import (
	"go-base/internal/adapter/http/route"
	"go-base/internal/domain/usecase"
	"go-base/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeExample(router *gin.Engine, db *gorm.DB, validate *validator.Validate) route.ExampleRoute {
	wire.Build(
		database.ProviderExampleRepository,
		usecase.ProviderExampleUseCase,
		route.ProviderExamplehRoute,
	)

	return nil
}

func InitializeHealth(router *gin.Engine) route.HealthRoute {
	wire.Build(
		route.ProviderHealthRoute,
	)

	return nil
}
