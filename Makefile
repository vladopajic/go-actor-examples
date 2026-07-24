GO ?= go
GOBIN ?= $$($(GO) env GOPATH)/bin
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint
GOLANGCI_LINT_VERSION ?= v2.0.1 # LINT_VERSION: update version in other places

.PHONY: install-golangcilint
install-golangcilint:
	test -f $(GOLANGCI_LINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

# Runs lint on entire repo
.PHONY: lint
lint: install-golangcilint
	$(GOLANGCI_LINT) run ./...

# Runs tests on entire repo
.PHONY: test
test: 
	go test --race -count=10 -failfast ./...

# Code tidy
.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...


# Runs example
.PHONY: run
run:
	go run ./cmd/... -example=$(filter-out $@,$(MAKECMDGOALS))

%:
	@:
