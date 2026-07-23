.PHONY: build
build:
	go build ./cmd/*/

.PHONY: run
run:
	go run ./cmd/nms-planner-cli/main.go