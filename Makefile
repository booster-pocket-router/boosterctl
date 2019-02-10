VERSION          := $(shell git describe --tags --always --dirty="-dev")
COMMIT           := $(shell git rev-parse --short HEAD)
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.version=$(VERSION)" -X "main.commit=$(COMMIT)" -X "main.buildTime=$(DATE)"'

#V := 1 # Verbose
Q := $(if $V,,@)

allpackages = $(shell ( cd $(CURDIR) && go list ./... ))
gofiles = $(shell ( cd $(CURDIR) && find . -iname \*.go ))

arch = "$(if $(GOARCH),_$(GOARCH)/,/)"
bind = "$(CURDIR)/bin/$(GOOS)$(arch)"
go = $(env GO111MODULE=on go)

.PHONY: all
all: cli

.PHONY: cli
cli:
	$Q go build $(if $V,-v) -o $(bind)/boosterctl $(VERSION_FLAGS) $(CURDIR)/main.go

.PHONY: clean
clean:
	$Q rm -rf $(CURDIR)/bin

.PHONY: format
format:
	$Q gofmt -s -w $(gofiles)
