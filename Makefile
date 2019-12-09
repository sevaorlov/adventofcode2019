PKGS = $(shell go list  -mod=vendor ./... | grep -v /test)

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -mod=vendor --race --cover $(PKGS)
.PHONY: test
