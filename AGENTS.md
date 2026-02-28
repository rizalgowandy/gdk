# Repository Guidelines

## Project Structure & Module Organization
This repository is a Go module (`github.com/rizalgowandy/gdk`) with reusable utilities under `pkg/`.  
Each package is self-contained (for example: `pkg/logx`, `pkg/httpx`, `pkg/errorx/v1`, `pkg/errorx/v2`) and usually includes its own `*_test.go` files.  
Root-level files include `loader.go`, module metadata (`go.mod`, `go.sum`), and task automation (`Taskfile.yml`).  
Supporting assets and docs live in `docs/` (including coverage artifacts and contributor docs), while CI and hooks are in `.github/`.

## Build, Test, and Development Commands
- `task tools`: installs local dev tooling (linters, generators, test tools) and sets up `.git/hooks/pre-commit`.
- `task generate`: runs `go generate ./...`.
- `task analysis`: runs `golangci-lint` on new changes (`--new-from-rev HEAD~`).
- `task unit_tests`: runs race-enabled verbose tests via `gotestsum`.
- `task build`: validates buildability by compiling a temporary binary.
- `go test ./... -coverprofile=docs/coverage.out`: generates a coverage profile (matches CI coverage workflow).

## Coding Style & Naming Conventions
Use standard Go formatting and imports; this repo enforces `gofmt` and `goimports` via `golangci-lint`.  
Pre-commit also applies `golines` to staged `.go` files.  
Follow idiomatic Go naming: lowercase package names, exported identifiers in `CamelCase`, tests in `*_test.go`.  
Keep package APIs focused; prefer adding new utilities to existing domain packages under `pkg/` rather than creating broad shared files.

## Testing Guidelines
Write table-driven tests where practical and keep tests close to implementation files in the same package directory.  
Use descriptive test names like `TestParseToken_InvalidSignature`.  
Run `task unit_tests` before opening a PR; for targeted checks, use `go test ./pkg/<name> -v -race`.

## Commit & Pull Request Guidelines
Recent history favors short, imperative commit subjects (for example, `Add Echo Custom Binder`) and may append PR refs like `(#50)`.  
Keep commits scoped to one logical change.  
For PRs targeting `main`, include:
- clear summary of behavior changes,
- linked issue/context,
- updated docs when interfaces/configs change (see `docs/CONTRIBUTING.md`),
- passing CI (`build`, tests, coverage).
