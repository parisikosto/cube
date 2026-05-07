APP_MODULE  := github.com/parisikosto/cube
APP_VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "development")
GIT_COMMIT  := $(shell git rev-list -1 HEAD 2>/dev/null || echo "unknown")

BINARY_NAME          := cube
BINARY_NAME_UBUNTU   := cube-linux-amd64
BINARY_NAME_MACOS    := cube-darwin-amd64
BINARY_NAME_RASPBIAN := cube-linux-arm

GOOS_LINUX  := linux
GOOS_DARWIN := darwin

GOARCH_AMD64 := amd64
GOARCH_ARM   := arm
GOARM_5      := 5

TARGET_OS_UBUNTU   := Ubuntu 24.04.4 LTS (linux/amd64)
TARGET_OS_MACOS    := macOS (darwin/amd64)
TARGET_OS_RASPBIAN := Raspbian GNU/Linux 10 (linux/arm)

DEFAULT_LDFLAGS := \
	-X '$(APP_MODULE)/build.Version=$(APP_VERSION)' \
	-X '$(APP_MODULE)/build.User=$(shell id -u -n)' \
	-X '$(APP_MODULE)/build.Time=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")' \
	-X '$(APP_MODULE)/build.GitCommit=$(GIT_COMMIT)'

.PHONY: build build-ubuntu build-macos build-raspbian build-all clean

build:
	@echo " > Building $(BINARY_NAME) $(APP_VERSION)..."
	@go build -ldflags="$(DEFAULT_LDFLAGS)" -o $(BINARY_NAME) main.go

build-ubuntu:
	@echo " > Building $(BINARY_NAME_UBUNTU) for $(TARGET_OS_UBUNTU)..."
	@env GOOS=$(GOOS_LINUX) GOARCH=$(GOARCH_AMD64) go build \
		-ldflags="$(DEFAULT_LDFLAGS) -X '$(APP_MODULE)/build.TargetOS=$(TARGET_OS_UBUNTU)'" \
		-o $(BINARY_NAME_UBUNTU) main.go

build-macos:
	@echo " > Building $(BINARY_NAME_MACOS) for $(TARGET_OS_MACOS)..."
	@env GOOS=$(GOOS_DARWIN) GOARCH=$(GOARCH_AMD64) go build \
		-ldflags="$(DEFAULT_LDFLAGS) -X '$(APP_MODULE)/build.TargetOS=$(TARGET_OS_MACOS)'" \
		-o $(BINARY_NAME_MACOS) main.go

build-raspbian:
	@echo " > Building $(BINARY_NAME_RASPBIAN) for $(TARGET_OS_RASPBIAN)..."
	@env GOOS=$(GOOS_LINUX) GOARCH=$(GOARCH_ARM) GOARM=$(GOARM_5) go build \
		-ldflags="$(DEFAULT_LDFLAGS) -X '$(APP_MODULE)/build.TargetOS=$(TARGET_OS_RASPBIAN)'" \
		-o $(BINARY_NAME_RASPBIAN) main.go

build-all: build-ubuntu build-macos build-raspbian

clean:
	@echo " > Removing binaries..."
	@rm -f $(BINARY_NAME) $(BINARY_NAME_UBUNTU) $(BINARY_NAME_MACOS) $(BINARY_NAME_RASPBIAN)
