APP_NAME := pufferstarter_cli
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

TARGETS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64

.PHONY: all clean

all: $(TARGETS)

$(TARGETS):
	$(eval OS := $(word 1,$(subst /, ,$@)))
	$(eval ARCH := $(word 2,$(subst /, ,$@)))
	$(eval EXT := $(if $(filter windows,$(OS)),.exe,))
	GOOS=$(OS) GOARCH=$(ARCH) go build $(LDFLAGS) \
		-o dist/$(APP_NAME)-$(OS)-$(ARCH)$(EXT) .

clean:
	rm -rf dist/

.PHONY: checksums
checksums:
	cd dist && sha256sum * > SHA256SUMS
