VERSION := v0.1.1

.PHONY: install
install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/juicer
