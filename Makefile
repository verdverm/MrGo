
# Default target executed when no arguments are given to make.
default_target: dev
.PHONY : default_target



#=============================================================================
# Main build rules
#=============================================================================

dev: directories head
.PHONY : dev

all: directories head node
.PHONY : all

clean:
	@rm -f bin/*
.PHONY : clean



#=============================================================================
# sub-rules for building
#=============================================================================

directories:
	@mkdir -p bin
.PHONY : directories

head:
	@echo "[building]  head"
	@go build -o head main/head.go
	@mv head bin/head
.PHONY : head

node:
	@echo "[building]  node"
	@go build -o node main/node.go
	@mv node bin/node
.PHONY : node


