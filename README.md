# Go Base SweetEscape Clean Architecture

This is a boilerplate for a Go SweetEscape project implementing Clean Architecture using the Gin web framework.

## Directory Structure

```
📦 project-root
├── 📂 cmd                 # Application entry point
├── 📂 config              # Application configuration
├── 📂 internal            # Core business logic
│   ├── 📂 domain          # Entities and Repository interface
│   ├── 📂 usecase         # Use cases (application logic)
│   ├── 📂 infrastructure  # External components (DB, HTTP, Middleware, Logger, HTTP handler (controller))
├── 📂 pkg                 # Helper utilities and additional libraries
├── go.mod                 # Module dependencies
└── go.sum                 # Dependency checksums
```

## Technologies Used

- [Go 1.24.0](https://golang.org/)
- [Gin (HTTP framework)](https://gin-gonic.com/)
- [GORM (ORM)](https://gorm.io/)
- [Viper (Application Configuration)](https://github.com/spf13/viper)
- [Zap (Logging)](https://github.com/uber-go/zap)
- [Air (Live Reload)](https://github.com/air-verse/air)

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

5. Run the application:
   ```sh
   air
   ```

## API Endpoint

| Method | Endpoint  | Deskripsi        |
| ------ | --------- | ---------------- |
| GET    | `/health` | Check API status |
