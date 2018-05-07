# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SOLC=solc
BINARY_NAME=mybinary

CONTRACTS=hello-world

all: contracts

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf *-target
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

contracts:  
	    @$(foreach CONTRACT,$(CONTRACTS),$(SOLC) -o $(CONTRACT)-target --bin --abi $(CONTRACT).sol;)

deps:
	$(GOGET) ./...
