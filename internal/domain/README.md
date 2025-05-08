# Domain Layer

This folder contains the core domain logic of the application. According to Clean Architecture, this is the **heart of the system**, containing business rules that are **independent** of frameworks, databases, and delivery mechanisms.

## Structure

```
internal/
└── domain/
    ├── entity/
    ├── repository/
    └── usecase/
```

### 📁 entity

Contains **enterprise entities** and **value objects** that encapsulate the core business rules.

- Entities represent key domain models (e.g., `User`, `Order`, `Product`).
- They are **framework-agnostic** and should contain only business logic.

> Keep entities simple, testable, and free from external dependencies.

---

### 📁 repository

Defines **interfaces (contracts)** for data access operations needed by the domain.

- Interfaces should only describe **what** needs to be done, not **how**.
- Implementation details (e.g., SQL, API calls) should live in the infrastructure layer.

Example:

```go
type (
	OrderRepository interface {
		FindByOrderCode(code string) (*entity.Order, error)
	}

	OrderRepositoryImpl struct {
		db       *gorm.DB
		validate *validator.Validate
	}
)

func ProviderOrderRepository(db *gorm.DB, validate *validator.Validate) OrderRepository {
	return &OrderRepositoryImpl{db, validate}
}

func (r *OrderRepositoryImpl) FindByOrderCode(code string) (*entity.Order, error) {
	if len(code) == 0 {
		return nil, errors.New("order code is required")
	}

	var order entity.Order
	err := r.db.Where("code = ?", code).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}
```

---

### 📁 usecase

Contains **application-specific business rules**.

- Coordinates actions between entities and repositories.
- Encapsulates the **application flow** or **user actions** (e.g., "RegisterUser", "PlaceOrder").

Use cases:

- Should not depend on external packages.
- Should **only depend on interfaces** from the `repository` package and entities from the `entity` package.

---

## Guiding Principles

- No dependencies on external frameworks or tools.
- Domain code should be **pure Go** and **unit-testable**.
- Easy to port or refactor without affecting external layers (e.g., HTTP, DB).

---

## 📌 Best Practices

- Name interfaces with meaningful business context (`UserRepository`, not `UserDBHandler`).
- Favor **dependency inversion**: the domain defines the interfaces, and other layers implement them.
- Avoid leaking HTTP/database concepts into the domain (no JSON tags, SQL annotations, etc.).

---

## Related Layers

- Infrastructure implementations (e.g., database access) are located in `internal/infrastructure/database`.
- HTTP handlers and routes live in `internal/adapter/http`.
