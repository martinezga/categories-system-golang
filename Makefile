GO = go
GOLANGCILINT = golangci-lint

run: ## run the API server
	@$(GO) run main.go

tools:  ## Install dev-tools
	@$(GO) get github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:  ## Run lint tools
	@$(GOLANGCILINT) run
