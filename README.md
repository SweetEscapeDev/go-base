# Go Base SweetEscape Clean Architecture

This is a boilerplate for a Go SweetEscape project implementing Clean Architecture using the Gin web framework.

## Directory Structure

```
📦 project-root
├── 📂 cmd                       # Application entry point (main package)
├── 📂 config                    # Application configuration (env, config loader)
├── 📂 internal                  # Core business logic (not exposed outside the module)
│   ├── 📂 domain                # Domain layer (enterprise business rules)
│   │   ├── 📂 entity            # Entity definitions and value objects
│   │   ├── 📂 repository        # Repository interfaces (data access contracts)
│   │   └── 📂 usecase           # Use case interfaces (application business rules)
│   ├── 📂 adapter               # Interface adapters (input/output handlers)
│   │   └── 📂 http              # HTTP layer (web handlers and controllers)
│   │       ├── 📂 dto           # Data Transfer Objects for request/response
│   │       └── 📂 route         # HTTP route definitions
│   └── 📂 infrastructure        # External systems and frameworks (framework & driver layer)
│       ├── 📂 database          # Repository implementations (database access)
│       ├── 📂 initializer       # Component initializers (DB, DI, config)
│       └── 📂 middleware        # HTTP middleware (logging, auth, etc.)
├── 📂 pkg                       # Shared utilities and helper functions
├── 📂 script                    # Custom scripts (migrations, seeders, etc.)
├── 📂 test                      # Test helpers, mocks, and fixtures
├── go.mod                       # Go module definition and dependencies
└── go.sum                       # Dependency checksums for verification
```

## Technologies Used

- [Go 1.24.0](https://golang.org/)
- [Gin (HTTP framework)](https://gin-gonic.com/)
- [GORM (ORM)](https://gorm.io/)
- [Viper (Application Configuration)](https://github.com/spf13/viper)
- [Zap (Logging)](https://github.com/uber-go/zap)
- [Air (Live Reload)](https://github.com/air-verse/air)
- [Wire (Depedency Injection Library)](https://github.com/google/wire)

## Installation and Running the Application

1. Clone this repository:

   ```sh
   git clone https://github.com/SweetEscapeDev/go-base.git
   cd go-base
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Configure environment:

   ```sh
   cp .env.example .env
   ```

4. Install air:

   ```sh
   go install github.com/air-verse/air@latest
   ```

5. Install wire:

   ```sh
   go install github.com/google/wire/cmd/wire@latest
   ```

6. Generate Wire:

   ```sh
   sh script/wire.sh
   ```

7. Run the application:
   ```sh
   air
   ```

## API Endpoint

| Method | Endpoint  | Deskripsi        |
| ------ | --------- | ---------------- |
| GET    | `/health` | Check API status |

## SCRIPT

1. wire.sh: generate wire_gen.go
