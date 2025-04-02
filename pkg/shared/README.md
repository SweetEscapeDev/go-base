# Pkg Layer

The `pkg` layer contains utility functions, helpers, and shared libraries that can be used throughout the application. These components are independent and reusable across different parts of the project.

## Responsibilities

- Provide utility functions that do not belong to a specific domain.
- Contain reusable helper packages for logging, error handling, and formatting.
- Keep general-purpose implementations separate from business logic.

## Directory Structure

```
📂 pkg
├── 📂 logger           # Shared logging utilities
│   ├── logger.go      # Zap logger setup
├── 📂 errors           # Custom error handling
│   ├── errors.go      # Standardized error responses
├── 📂 utils            # General utility functions
│   ├── utils.go       # Helper functions (e.g., string manipulation, validation)
└── ...
```

## Example Logger Implementation

```go
package logger

import (
    "go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
    logger, _ := zap.NewProduction()
    Log = logger
}
```

## Best Practices

- Keep functions generic and not tied to a specific use case.
- Avoid business logic in utility functions.
- Make sure utilities are well-documented and easily reusable.

This helps maintain a clean separation of concerns and ensures utilities can be leveraged across different modules.
