LOCAL_BIN = $(CURDIR)/bin

run:
	docker compose up -d
	go run ./cmd/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v3@v3.5.5

generate-mocks:
	GOBIN=$(LOCAL_BIN)/mockery


.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out