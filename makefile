include .env

main_package_path = ./cmd/main.go
binary_name = f1-api.exe
build_dir = bin

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  tidy       - tidy modfiles and format .go files"
	@echo "  build      - build the application"
	@echo "  run        - run the application"
	@echo "  run/live   - run the application with reloading on file changes"

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && set /p ans= && if /i not "%ans%"=="y" exit 1

.PHONY: no-dirty
no-dirty:
	@git status --porcelain >nul 2>&1 || (echo "Working directory is dirty!" && exit 1)

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	@if not exist $(build_dir) mkdir $(build_dir)
	go build -o=$(build_dir)\$(binary_name) $(main_package_path)

## run: run the application
.PHONY: run
run: build
	$(build_dir)\$(binary_name)

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	air \
		--build.cmd "make build" --build.bin "$(build_dir)\$(binary_name)" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"