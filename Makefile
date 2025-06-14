.DEFAULT_GOAL:= build

.PHONY: run
run:
	go run ./cmd/psps/main.go

.PHONY: fmt vet build
build:
	go build -o ./bin/psps ./cmd/psps/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test --cover ./...
