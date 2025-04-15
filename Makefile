.PHONY: format imports tidy vet staticcheck lint build precommit test


format:
	@echo "Formatting code..."
	gofmt -s -w ${GOFILES_NOT_NODE}


imports: install-goimport
	@echo "Running goimports..."
	goimports -l -w ${GOFILES_NOT_NODE}


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