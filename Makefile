BINDIR ?= $(CURDIR)/bin

all: exporter

exporter: $(BINDIR)/node-exporter

prepare:
	go mod tidy

clean:
	rm -r $(BINDIR)

$(BINDIR)/node-exporter: prepare
	CGO_ENABLED=0 go build -o $@ ./cmd/main.go

.PHONY: all prepare clean

