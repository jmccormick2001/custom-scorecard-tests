IMAGE = quay.io/jemccorm/custom-scorecard-tests
SHELL = /bin/bash

all: build

clean: ## Clean up the build artifacts
	rm -rf build

build:
	go build internal/tests/tests.go
image-build:
	go build -o images/custom-scorecard-tests/custom-scorecard-tests images/custom-scorecard-tests/cmd/test/main.go
	cd images/custom-scorecard-tests && docker build -t $(IMAGE):dev .
image-push:
	./hack/image/push-image-tags.sh $(IMAGE):dev $(IMAGE)-$(shell go env GOARCH)
