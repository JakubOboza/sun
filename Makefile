GO=go

.PHONY: default
default: test ;

install: 
	$(GO) get github.com/stianeikeland/go-rpio

test: 
	$(GO) test -v pot/*
