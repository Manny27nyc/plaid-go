# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=1.2.0

.PHONY: test
test:
	go test -v ./...
