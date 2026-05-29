VERSION := v0.1.6

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/codist
