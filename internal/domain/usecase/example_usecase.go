package usecase

import (
	"go-base/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type (
	ExampleUseCase interface {
		ExampleCase(c *gin.Context) (bool, error)
	}

	exampleUseCaseImpl struct {
		ExampleRepository repository.ExampleRepository
	}
)

func (e *exampleUseCaseImpl) ExampleCase(c *gin.Context) (bool, error) {
	return true, nil
}

func ProviderExampleUseCase(
	exampleRepository repository.ExampleRepository,
) ExampleUseCase {
	return &exampleUseCaseImpl{
		exampleRepository,
	}
}
