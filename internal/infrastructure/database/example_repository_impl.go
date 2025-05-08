package database

import (
	"go-base/internal/domain/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type exampleRepositoryImpl struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (e *exampleRepositoryImpl) ExampleMethod() (bool, error) {
	return true, nil
}

func ProviderExampleRepository(db *gorm.DB, validate *validator.Validate) repository.ExampleRepository {
	return &exampleRepositoryImpl{db, validate}
}
