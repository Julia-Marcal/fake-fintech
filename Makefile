# Makefile for fake-fintech — single definition

GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
BINARY := cmd/publish-message/publish-message

# Defaults (override on the command line)
ACTION ?= FETCH_PRICE
MARKET ?= crypto
SYMBOL ?= BTC
URL ?= $(RABBITMQ_URL)

.PHONY: help build-publish publish publish-binary clean

help:
	@echo "Makefile targets:"
	@echo "  make publish           -> run publisher with 'go run' (uses ACTION, MARKET, SYMBOL, URL)"
	@echo "  make publish-binary    -> build and run the publisher binary"
	@echo "  make build-publish     -> build the publisher binary only"
	@echo "  make clean             -> remove built binary"
	@echo "Examples:"
	@echo "  make publish ACTION=FETCH_PRICE MARKET=crypto SYMBOL=BTC URL=amqp://guest:guest@localhost:5672/"
	@echo "  RABBITMQ_URL=amqp://guest:guest@localhost:5672/ make publish"

# Build the binary

build-publish:
	$(GOBUILD) -o $(BINARY) ./cmd/publish-message

# Run with 'go run' (no binary produced)
publish:
	$(if $(URL),$(GORUN) ./cmd/publish-message --action "$(ACTION)" --market "$(MARKET)" --symbol "$(SYMBOL)" --url "$(URL)",$(error ERROR: RabbitMQ URL not set - provide URL or set RABBITMQ_URL env var))

# Build then run the binary
publish-binary: build-publish
	@if [ -z "$(URL)" ]; then \
		echo "ERROR: RabbitMQ URL not set. Provide URL or set RABBITMQ_URL env var."; exit 1; \
	fi
	$(BINARY) --action "$(ACTION)" --market "$(MARKET)" --symbol "$(SYMBOL)" --url "$(URL)"

clean:
	-@rm -f $(BINARY)
