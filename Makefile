BIN_DIR ?= $(CURDIR)/bin

main: exporter

exporter: $(BIN_DIR)/exporter

prepare:
	go mod tidy

clean:
	rm -r $(BIN_DIR)

$(BIN_DIR)/exporter: clean prepare
	CGO_ENABLED=0 go build -o $@ ./cmd/main.go

.PHONY: main prepare clean

