GO = go
GOLANGCILINT = golangci-lint

# Load envs from .env file if it exists
# If variable names or values have spaces or special characters, you need to quote them
-include .env
export

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_/-]+:.*?##/ { printf "\033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

install: ## install dependencies
	@$(GO) mod download

run: ## run the API server
	@$(GO) run main.go

docker/up:  ## Build and run docker apps
	@docker compose up --build

docker/down:  ## Stop docker apps
	@docker compose down

tools:  ## Install dev-tools
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@$(GO) install github.com/swaggo/swag/cmd/swag@latest

lint:  ## Run lint tools
	@$(GOLANGCILINT) run

swagger:
	swag init -generalInfo api/api.go -output ./docs/swagger
