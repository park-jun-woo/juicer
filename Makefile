VERSION := v0.1.7

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/codist
