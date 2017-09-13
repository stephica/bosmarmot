#
# |_  _  _ _ _  _  _ _ _  _ _|_
# |_)(_)_\| | |(_|| | | |(_) |
#
# Bosmarmot Makefile
#
# Requires go version 1.8 or later.
#

SHELL := /bin/bash
GOFILES_NOVENDOR := $(shell find . -type f -name '*.go' -not -path "**/vendor/*")
GOPACKAGES_NOVENDOR := $(shell go list ./...)
OS_ARCHS := "linux/arm linux/386 linux/amd64 darwin/386 darwin/amd64 windows/386 windows/amd64"

# Install dependencies and also clear out vendor (we should do this in CI)

# Print version
.PHONY: version
version:
	@go run ./release/cmd/version/main.go

# Run goimports (also checks formatting) first display output first, then check for success
.PHONY: check
check:
	@goimports -l -d ${GOFILES_NOVENDOR}
	@goimports -l ${GOFILES_NOVENDOR} | read && echo && \
	echo "Your marmot has found a problem with the formatting style of the code."\
	 1>&2 && exit 1 || true

# Just fix it
.PHONY: fix
fix:
	@goimports -l -w ${GOFILES_NOVENDOR}

# Run tests
.PHONY:	test
test: check
	@go test ${GOPACKAGES_NOVENDOR}

# Run tests for development (noisy)
.PHONY:	test_dev
test_dev:
	@go test -v ${GOPACKAGES_NOVENDOR}

# Run tests including integration tests
.PHONY:	test_integration
test_integration: check
	@go test -tags integration ${GOPACKAGES_NOVENDOR}


.PHONY: build_bin
build_bin:
	@go build -o bin/bos ./monax/cmd/bos
	@go build -o bin/monax-keys ./keys/cmd/monax-keys

# Build all the things
.PHONY: build
build:	build_bin

# Build binaries for all architectures
.PHONY: build_dist
build_dist:
	@goreleaser --rm-dist --skip-publish --skip-validate

# Generate full changelog of all release notes
changelog.md: ./release/release.go
	@go run ./release/cmd/changelog/main.go > changelog.md

# Generated release notes for this version
notes.md: ./release/release.go
	@go run ./release/cmd/notes/main.go > notes.md

# Do all available tests and checks then build
.PHONY: build_ci
build_ci: check test build

# Tag the current HEAD commit with the current release defined in
# ./release/release.go
.PHONY: tag_release
tag_release: test check changelog.md build_bin
	@scripts/tag_release.sh

# If the checked out commit is tagged with a version then release to github
.PHONY: release
release: notes.md
	@scripts/release.sh
