VERSION := v0.1.0
BINARY  := huma

.PHONY: build install clean

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY) .

install:
	go install -ldflags "-X main.Version=$(VERSION)" .

clean:
	rm -f $(BINARY)
