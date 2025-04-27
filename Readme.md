# Go-Logger

A lightweight, extensible logging library for Go built on `slog`. Supports multi-output logging (console + file), colored terminal output via `tint`, and dynamic log level control.

[![Go Reference](https://pkg.go.dev/badge/github.com/VolkHackVH/go-logger.svg)](https://pkg.go.dev/github.com/VolkHackVH/go-logger)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/VolkHackVH/go-logger/blob/main/LICENSE)

## 📍 Installation

```text
go get github.com/VolkHackVH/go-logger
```

## 🌟 Features

- **Log level:** `Log`, `Debug`, `Warn`, `Error`
- **Level control:** Toggle between `DEBUG` and `ERROR` modes at initialization
- **Dual output**: Simultaneous logging to console (with colors) and files
- **Structured logging**: Compatible with `slog.Attr` and context values
- **File rotation**: Helper function for log file creation
- **Zero-dependency**: Only uses standard library + `tint` for colors

## ⚡ Example

```go
package main

import (
    "github.com/VolkHackVH/go-logger"
)

func main() {
    /* Enable debug logging (false would show only errors) */
    logger.NewLogger(true, "./Logger/Logger.log")

    logger.Log("Initialized...")
    logger.Debug("Connected to Database...")
    logger.Warn("Detected...")
    err := logger.Error("Connection failed: %v", err) // returns error
}
```
