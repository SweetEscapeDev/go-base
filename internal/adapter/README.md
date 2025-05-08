# Adapter Layer

This folder contains the **interface adapters** responsible for translating data between the domain layer and external delivery mechanisms (e.g., HTTP, CLI, gRPC). It acts as a bridge between the core logic (domain/usecase) and the outside world.

## Structure

```
internal/
└── adapter/
    └── http/
        ├── dto/
        └── route/
```

### 📁 http

Responsible for handling HTTP requests and responses. It integrates with the domain's use cases through controllers/handlers.

#### 📁 dto (Data Transfer Object)

- Defines request and response payload formats for the HTTP layer.
- Translates between external JSON structures and internal entities.
- Should not contain business logic.

Example:

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

#### 📁 route

- Defines HTTP endpoints and maps them to corresponding handler functions.
- Routes are grouped logically (e.g., user, product, order).
- Initializes handlers with the required use case dependencies (often injected).

---

## Responsibilities

- Accept input from the external world (e.g., user requests via HTTP).
- Validate, format, and map input into domain-level entities or use cases.
- Call the appropriate use case(s) and return formatted output.

---

## Guiding Principles

- This layer should depend **on domain/usecase**, not the other way around.
- Avoid putting business logic here—delegate to the use case layer.
- Each adapter should be focused on a specific interface (e.g., `http`, `grpc`, `cli`).

---

## Related Layers

- Business logic is handled in `internal/domain/usecase`.
- External frameworks and middleware are initialized in `internal/infrastructure`.
