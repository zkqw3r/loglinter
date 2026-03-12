# loglinter

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://go.dev/)
[![golangci-lint](https://img.shields.io/badge/golangci--lint-v2.11.3-blue)](https://golangci-lint.run/)
[![zap](https://img.shields.io/badge/zap-supported-orange)](https://github.com/uber-go/zap)
[![slog](https://img.shields.io/badge/slog-supported-green)](https://pkg.go.dev/log/slog)


Go linter for checking log messages, compatible with golangci-lint

## Verification Rules

| Rule | Example ❌ | Example ✅ |
|---|---|---|
| Lowercase letter at the beginning | `“Server started”` | `“server started”` |
| English only | `“запуск сервера”` | `“server started”` |
| No special characters or emojis | `“starting! 🚀”` | `“starting”` |
| No sensitive data | `“user password: 123”` | `“user authenticated”` |
## Supported loggers

- `go.uber.org/zap` — `Logger` and `SugaredLogger`
- `log/slog` — package functions and methods `*slog.Logger`

### Auto-fix

Run with `--fix` to automatically correct violations:

```bash
./custom-gcl run --fix ./...
```

Fixable rules: uppercase → lowercase, removes special symbols and emojis.


## How it works

The linter analyzes calls to `go.uber.org/zap` and `log/slog` logging functions,
extracts the message argument and validates it against the configured rules.
All checks are configurable via `.golangci.yml`. Rules support auto-fix —
run with `--fix` flag to apply corrections automatically.

## Installation and launch

### Requirements
- Go 1.22+
- golangci-lint v2.11.3+

### Building a custom binary

```bash
golangci-lint custom
```

Creates `./custom-gcl` — golangci-lint with the plugin connected.

### Running

```bash
./custom-gcl run ./...
```

## Configuration

In `.golangci.yml`, you can configure which checks are enabled and the list of sensitive words:

```yaml
version: "2"

linters:
  default: none
  enable:
    - loglinter
  settings:
    custom:
      loglinter:
        type: "module"
        description: "Checks log messages style"
        settings:
          check_uppercase: true
          check_language: true
          check_special_symbols: true
          check_sensitive: true
          sensitive:
            - "password"
            - "token"
            - "secret"
            - "api_key"
```

## Example

```
❯ ./custom-gcl run test_sample.go

test_sample.go:8:14: Log messages should start with a lowercase letter (loglinter)
        sugar.Infow("Server started", "port", 8080)
                    ^
test_sample.go:9:14: Log messages must be in English only (loglinter)
        sugar.Infow("запуск сервера", "port", 8080)
                    ^
test_sample.go:10:14: Log messages should not contain special characters or emojis (loglinter)
        sugar.Infow("starting!", "port", 8080)
                    ^
3 issues:
* loglinter: 3
```

## Tests

```bash
go test ./...
```

## Project structure

\```
loglinter/
├── cmd/logchecker/main.go       # standalone binary entry point
├── pkg/analyzer/
│   ├── analyzer.go              # analyzer definition and run function
│   ├── logmessage.go            # log detection and message extraction
│   ├── checker.go               # rule checker functions
│   ├── checker_test.go          # unit tests for each checker
│   ├── analyzer_test.go         # integration test via analysistest
│   └── config.go                # Config struct, Result, method maps
│   └── testdata/src/
│       ├── go.uber.org/zap/     # zap stub for tests
│       └── sample/sample.go     # test cases with // want comments
├── plugin/plugin.go             # golangci-lint module plugin
├── .golangci.yml                # linter configuration example
├── .custom-gcl.yml              # golangci-lint custom build config
└── .github/workflows/ci.yml     # CI/CD
\```
