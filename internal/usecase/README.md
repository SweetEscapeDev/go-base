# Usecase Layer

The `usecase` layer is responsible for implementing business logic and coordinating interactions between repositories and external services. This layer acts as an intermediary between the HTTP handlers (controllers) and the repository layer.

## Responsibilities

- Implement application-specific business logic.
- Communicate with the repository layer for data persistence.
- Ensure data integrity and business rule enforcement.
- Provide an abstraction layer between the HTTP handlers and repositories.

## Directory Structure

```
📂 usecase
├── 📂 user               # Business logic for user-related operations
│   ├── user_usecase.go   # Implementation of user usecase
│   ├── dto.go            # Data transfer objects (DTOs) for users
│   ├── interface.go      # Interface defining the user usecase methods
├── 📂 order              # Business logic for order-related operations
│   ├── order_usecase.go  # Implementation of order usecase
│   ├── dto.go            # Data transfer objects (DTOs) for orders
│   ├── interface.go      # Interface defining the order usecase methods
└── ...
```

## Example Usecase Implementation

```go
package user

type UserUsecase struct {
    repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
    return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetUserByID(id int) (*UserDTO, error) {
    user, err := uc.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    return &UserDTO{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}
```

## Testing

To ensure the correctness of the business logic, unit tests should be implemented for each use case.

Example test using Go’s built-in testing package:

```go
package user_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
    mockRepo := new(MockUserRepository)
    usecase := NewUserUsecase(mockRepo)

    mockRepo.On("FindByID", 1).Return(&User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)

    user, err := usecase.GetUserByID(1)
    assert.NoError(t, err)
    assert.Equal(t, "John Doe", user.Name)
}
```

## Best Practices

- Keep business logic inside the usecase layer and avoid direct repository calls in controllers.
- Use interfaces to decouple dependencies.
- Implement unit tests with mocks to verify logic independently.

This ensures the separation of concerns and maintainability in the Clean Architecture pattern.
