GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=play-echo

run: $(GORUN)
build: $(GOBUILD) -o $(BINARY_NAME) -v
test: $(GOTEST) -v ./...
clean: $(GOCLEAN) && rm -f $(BINARY_NAME)
run: $(GOBUILD) -o $(BINARY_NAME) -v ./... && ./$(BINARY_NAME)