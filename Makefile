.DEFAULT_GOAL := help
.PHONY: test lint check format cover help

PACKAGES=./client ./model

test: ## Run tests
	go test ./client/... -timeout=1m -cover
	go test ./model/... -timeout=1m -cover

test-race: ## Run tests with -race. Note: expected to fail, but look for "DATA RACE" failures specifically
	go test ./client/... -timeout=2m -race
	go test ./model/... -timeout=2m -race

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	gometalinter --deadline=3m -j 2 --disable-all --tests --vendor \
		-E deadcode \
		-E errcheck \
		-E gas \
		-E goconst \
		-E gofmt \
		-E goimports \
		-E golint \
		-E ineffassign \
		-E interfacer \
		-E maligned \
		-E megacheck \
		-E misspell \
		-E nakedret \
		-E structcheck \
		-E unconvert \
		-E unparam \
		-E varcheck \
		-E vet \
		./...

check: lint ## Run tests and linters

cover: ## Runs tests on ./src/ with HTML code coverage
	@echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	go get -u github.com/alecthomas/gometalinter
	gometalinter --vendored-linters --install

format:  # Formats the code. Must have goimports installed (use make install-linters).
	# This sorts imports by [stdlib, 3rdpart, mdllife/teller]
	goimports -w -local github.com/modeneis/waves-go-client ./client
	goimports -w -local github.com/modeneis/waves-go-client ./model
	# This performs code simplifications
	gofmt -s -w ./client
	gofmt -s -w ./model

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'