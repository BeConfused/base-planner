.PHONY: init
init:
	@go install ./cmd/*/

.PHONY: build
build:
	@go build ./cmd/*/

.PHONY: run
run:
	@go run ./cmd/base-planner/main.go

.PHONY: lint
lint:
	@golangci-lint run