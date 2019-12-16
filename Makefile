PKGS = $(shell go list  -mod=vendor ./... | grep -v /test)

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -mod=vendor --race --cover $(PKGS)
.PHONY: test

# works only for MacOS, remove an argument after -i for it to work in other OS
build_day:
	mkdir day$(DAY)
	cp -R day_skeleton/* day$(DAY)
	find ./day$(DAY) -type f -print0 | xargs -0 sed -i '' 's/day_skeleton/day$(DAY)/g'
.PHONY: build_day
