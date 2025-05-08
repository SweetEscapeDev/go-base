# Infrastructure Layer

This folder contains implementations of external systems and services used by the application. It serves as the **outermost layer** in Clean Architecture, where frameworks, drivers, and technical details live.

## Structure

```
internal/
└── infrastructure/
    ├── database/
    ├── external/
    ├── initializer/
    └── middleware/
```

### 📁 database

- Contains implementations of repository interfaces defined in the domain layer.
- Responsible for interacting with data sources (e.g., SQL databases, NoSQL, external APIs).
- Translates between domain entities and database models.

> This is where you implement interfaces like `UserRepository` using GORM, Ent, etc.

Example:

```go
type UserRepositorImpl struct {
    db *gorm.DB
}

func ProviderUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositorImpl{db}
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
    // DB query logic here
}
```

---

### 📁 external

- Contains integrations with external services and APIs (e.g., payment gateways, email providers, cloud services).
- Useful for wrapping third-party SDKs or creating abstractions for external interactions.
- Helps keep infrastructure-specific logic isolated from domain and use case layers.

---

### 📁 initializer

- Sets up and initializes infrastructure dependencies:
  - Database connections
  - Logger (e.g., Zap)
  - Background workers (e.g., Asynq)
  - Dependency injection wiring (e.g., with `google/wire`)

---

### 📁 middleware

- Contains custom HTTP middleware for:
  - Logging
  - Authentication & Authorization
  - Request tracing, CORS, panic recovery, etc.
- These are plugged into the HTTP router (e.g., Gin).

---

## Responsibilities

- Glue between external systems and the core application logic.
- Provide concrete implementations for interfaces declared in the domain/usecase layers.
- Handle technical concerns, not business rules.

---

## Guiding Principles

- This layer **depends on the domain**, not vice versa.
- Keep technical code here—avoid mixing domain logic.
- Easy to replace (e.g., switch from PostgreSQL to MongoDB) without affecting core business rules.

---

## Related Layers

- Domain interfaces are defined in `internal/domain/repository`.
- Use cases that consume infrastructure live in `internal/domain/usecase`.
- HTTP layer (routes, handlers) lives in `internal/adapter/http`.
