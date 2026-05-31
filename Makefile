VERSION := v0.1.11

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/codist
