GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
BUILD_VERSION := $(shell git log -1 --pretty=format:"%h (%ci)")
DIST_PATH := dist/$(GOOS)-$(GOARCH)
ZIP_NAME := strac-$(GOOS)-$(GOARCH).zip

build:
	rm -rf $(DISTPATH)

	go build \
		-ldflags "-X 'main.buildVersion=$(BUILD_VERSION)' \
			-extldflags '-static'" \
		-o $(DIST_PATH)/strac .

zip:
	rm -rf dist/$(ZIP_NAME)

	cd $(DIST_PATH) && zip -r ../$(ZIP_NAME) .

dist:
	rm -rf dist/

	GOOS=linux GOARCH=amd64 make build zip
	GOOS=darwin GOARCH=amd64 make build zip
	GOOS=windows GOARCH=amd64 make build zip

.PHONY: dist
