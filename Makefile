.PHONY: install-goimport install-staticcheck install-ginkgo install-devtools \
	format imports tidy vet staticcheck lint build precommit test


install-goimport:
	@if ! command -v goimports &> /dev/null; then \
		echo "goimports not found, installing..."; \
		go install golang.org/x/tools/cmd/goimports@v0.24.0; \
	fi

install-staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
		echo "staticcheck not found, installing..."; \
		go install honnef.co/go/tools/cmd/staticcheck@v0.4.7; \
	fi

install-ginkgo:
	@if ! command -v ginkgo &> /dev/null; then \
		echo "ginkgo not found, installing..."; \
		go install github.com/onsi/ginkgo/v2/ginkgo@v2.8.1; \
	fi

install-devtools: install-goimport install-staticcheck install-ginkgo


format:
	@echo "Formatting code..."
	gofmt -s -w zstd_compression/*.go


imports: install-goimport
	@echo "Running goimports..."
	goimports -l -w zstd_compression/*.go


tidy:
	@echo "Running go mod tidy..."
	go mod tidy


vet:
	@echo "Running go vet..."
	go vet ./...


staticcheck: install-staticcheck
	@echo "Running staticcheck..."
	staticcheck ./...


lint: format imports tidy vet staticcheck


build:
	@echo "Building..."
	go build ./...


precommit: lint test


test: install-ginkgo
	@echo "Running tests..."
	@ginkgo zstd_compression/


vendor:
	@echo "Vendoring..."
	go mod vendor