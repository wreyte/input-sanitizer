
.PHONY: tools lint test coverage cyclomatic clean tidy check-fmt check-deps vulnerable ci

TOOLS := honnef.co/go/tools/cmd/staticcheck@latest \
         golang.org/x/vuln/cmd/govulncheck@latest \
         github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
         github.com/fzipp/gocyclo/cmd/gocyclo@latest

tools:
	@echo "Installing/updating Go tools..."
	@for tool in $(TOOLS); do \
		go install $$tool; \
	done

lint: tools
	@echo "Linting..."
	@golangci-lint run ./...

test:
	@echo "Running tests with race detection..."
	@go test -v -race ./...

coverage:
	@echo "Running tests and generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@echo "Coverage report generated: coverage.out"

cyclomatic: tools
	@echo "Checking cyclomatic complexity..."
	@gocyclo -over 20 .

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin coverage.out

tidy:
	@echo "Running go mod tidy and go fmt (modifies files)..."
	@go mod tidy -v
	@go fmt ./...

check-fmt:
	@echo "Checking formatting..."
	@fmt_out=$$(gofmt -l .); \
		if [ -n "$$fmt_out" ]; then \
			echo "Unformatted files:"; \
			echo "$$fmt_out"; \
			exit 1; \
		fi

check-deps:
	@echo "Running go mod tidy -diff (checking for un-tidied modules)..."
	@go mod tidy -diff
	@echo "Verifying modules..."
	@go mod verify

vulnerable: tools
	@echo "Running go vulnerability check..."
	@govulncheck ./...

ci: clean tools check-deps check-fmt lint test coverage cyclomatic vulnerable
	@echo "All CI checks passed successfully!"

