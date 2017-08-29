DEP := $(shell command -v dep 2>/dev/null)
VERSION := $(shell git describe --tags 2> /dev/null || echo unknown)
IDENTIFIER := $(VERSION)-$(GOOS)-$(GOARCH)
CLONE_URL=github.com/ansrivas/gowindow
PKGS := $(shell cd $(GOPATH)/src/$(CLONE_URL); go list ./... | grep -v vendor)
BUILD_TIME=`date -u +%FT%T%z`
LDFLAGS="-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

.DEFAULT_GOAL := help

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:          ## Run all the tests
test:
	./test.sh

dep:           ## Get all the dependencies
dep:
	go get -u github.com/golang/dep/cmd/dep

vendor: Gopkg.toml
ifndef DEP
	make dep
endif
	dep ensure
	touch vendor

Gopkg.lock: dep Gopkg.toml
	dep ensure

clean:         ## Clean the application and remove all the docker containers.
clean:
	@go clean -i ./...

gowindow: vendor	clean
	go build -ldflags $(LDFLAGS) $(FLAGS) $(CLONE_URL)

crossbuild: vendor
	mkdir -p build/gowindow-$(IDENTIFIER)
	make gowindow FLAGS="-o build/gowindow-$(IDENTIFIER)/gowindow"
	@echo "Created release build: build/gowindow-$(IDENTIFIER).tar.gz"

release:       ## Create a release build.
release:
	make crossbuild GOOS=linux GOARCH=amd64
	make crossbuild GOOS=linux GOARCH=386
	make crossbuild GOOS=darwin GOARCH=amd64

bench:	       ## Benchmark the code.
bench:
	@go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

prof:          ## Run the profiler.
prof:	bench
	@go tool pprof cpu.prof

prof_svg:      ## Run the profiler and generate image.
prof_svg:	clean	bench
	@echo "Do you have graphviz installed? sudo apt-get install graphviz."
	@go tool pprof -svg gowindow.test cpu.prof > cpu.svg
