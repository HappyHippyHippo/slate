MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./... | grep -v dig))
TAGS	 = -tags="sqlite mysql postgre"
TESTPKGS = $(shell env GO111MODULE=on $(GO) list $(TAGS) ./... | grep -v dig)
BIN      = $(CURDIR)/_bin

GO      = go
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

export GO111MODULE=on

.PHONY: all
all: fmt lint critic tests

# Tools

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLINT = $(BIN)/revive
$(BIN)/revive: PACKAGE=github.com/mgechev/revive

GOCRITIC = $(BIN)/gocritic
$(BIN)/gocritic: PACKAGE=github.com/go-critic/go-critic/cmd/gocritic

GOCOV = $(BIN)/gocov
$(BIN)/gocov: PACKAGE=github.com/axw/gocov/...

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: PACKAGE=github.com/AlekSi/gocov-xml

GO2XUNIT = $(BIN)/go2xunit
$(BIN)/go2xunit: PACKAGE=github.com/tebeka/go2xunit

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt lint critic ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test $(TAGS) -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

test-xml: fmt lint critic | $(GO2XUNIT) ; $(info $(M) running xUnit tests…) @ ## Run tests with xUnit output
	$Q mkdir -p test
	$Q 2>&1 $(GO) test $(TAGS) -timeout $(TIMEOUT)s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input _test/tests.output -output _test/tests.xml

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML     = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML    = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOV) $(GOCOVXML)
test-coverage: COVERAGE_DIR := $(CURDIR)/_test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: fmt lint critic test-coverage-tools ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)
	$Q $(GO) test $(TAGS) \
		-coverpkg=$($(GO) list $(TESTPKGS)) \
		-covermode=$(COVERAGE_MODE) \
		-coverprofile="$(COVERAGE_PROFILE)" $(TESTPKGS)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running revive…) @ ## Run revive
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: critic
critic: | $(GOCRITIC) ; $(info $(M) running gocritic…) @ ## Run gocritic
	$Q $(GOCRITIC) check -enableAll $(PKGS)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt $(PKGS)

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
