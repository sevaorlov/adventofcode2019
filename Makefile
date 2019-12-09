PKGS = $(shell go list  -mod=vendor ./... | grep -v /test)

lint:
	golangci-lint run ./...
.PHONY: lint

run:
	go run main.go
.PHONY: run

test:
	go test -mod=vendor --race --cover $(PKGS)
.PHONY: test
