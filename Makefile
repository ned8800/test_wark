COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html
COVERIGNORE_FILE := cover_ignore.txt
BUILD_NAME := solution.a


UNAME := $(shell uname -s)
ifeq ($(UNAME), Linux)
    OPEN_CMD = xdg-open
else ifeq ($(UNAME), Darwin)
    OPEN_CMD = open
else
    OPEN_CMD = start
endif

TEST_PACKAGES := $(shell go list ./...)

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

build:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_NAME)$(if $(filter windows,$(GOOS)),.exe,) ./main.go


run: build
	@./$(BUILD_NAME)

test:
	@go test -coverprofile=$(COVERAGE_FILE) -covermode=atomic $(TEST_PACKAGES)

html: test
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	$(OPEN_CMD) $(COVERAGE_HTML)

coverage: test
	@go tool cover -func=$(COVERAGE_FILE) | grep total:

clean:
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

.PHONY: build run test html coverage clean
