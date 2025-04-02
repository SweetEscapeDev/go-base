# Infrastructure Layer

The `infrastructure` layer is responsible for handling external components such as databases, logging, and third-party integrations. It provides concrete implementations for interfaces defined in the `domain` layer.

## Responsibilities

- Implement repository interfaces for data persistence.
- Handle external services (e.g., caching, messaging, third-party APIs).
- Manage logging, configuration, and environment setup.

## Directory Structure

```
📂 infrastructure
├── 📂 db                # Database connection and repository implementations
│   ├── db.go           # Database connection setup
│   ├── user_repo.go    # User repository implementation
├── 📂 logger            # Logging setup using Zap
│   ├── logger.go       # Logger configuration
├── 📂 config            # Application configuration (Viper)
│   ├── config.go       # Load and manage configurations
└── ...
```

## Example Repository Implementation

```go
package db

import (
    "gorm.io/gorm"
    "myapp/internal/domain/user"
)

type UserRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindByID(id int) (*user.User, error) {
    var u user.User
    if err := r.db.First(&u, id).Error; err != nil {
        return nil, err
    }
    return &u, nil
}
```

## Best Practices

- Keep infrastructure concerns separate from business logic.
- Use dependency injection to make infrastructure components interchangeable.
- Ensure repositories only interact with database logic and do not contain business rules.

This separation ensures that the core business logic remains independent and easy to test.
