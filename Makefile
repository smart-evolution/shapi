GOCMD=go
GOLINT=golint
GOFMT=gofmt
MAKE=make
IMAGE_NAME="oszura/sh-dashboard-dev"

.DEFAULT_GOAL := all

.PHONY: install
install:
	$(shell cd /; $(GOCMD) get -u golang.org/x/lint/golint)
	$(GOCMD) mod vendor

.PHONY: all
all:
	$(GOCMD) build -mod=vendor -o smarthome

.PHONY: test
test:
	$(GOCMD) test -mod=vendor ./...

.PHONY: lint
lint:
	./scripts/gofmt_test.sh
	$(GOLINT) ./... | grep -v vendor/ && exit 1 || exit 0
	$(GOCMD) vet -mod=vendor ./... | grep -v vendor/ && exit 1 || exit 0

.PHONY: fix
fix:
	$(GOFMT) -w .

.PHONY: build-image
build-image:
	docker build --tag $(IMAGE_NAME) --file=./docker/Dockerfile .

.PHONY: run-container
run-container:
	docker run --network=docker_default -it -v $(PWD):/root/go/src/github.com/smart-evolution/smarthome \
	    -e SH_MONGO_URI=mongodb://172.18.0.2:27017 \
	    -e SH_MONGO_DB=smarthome \
	    -e SH_PANEL_PORT=3222 $(IMAGE_NAME)


.PHONY: version
version:
	git tag $(V)
	./scripts/changelog.sh
	go generate
	git add ./version.go || true
	git add ./docs/changelogs/CHANGELOG_$(V).md
	git commit --allow-empty -m "Build $(V)"
	git tag --delete $(V)
	git tag $(V)

.PHONY: help
help:
	@echo  '=================================='
	@echo  'Available tasks:'
	@echo  '=================================='
	@echo  '* Installation:'
	@echo  '- install         - Phony task that installs all required (client'
	@echo  '                    and server - sided) dependencies'
	@echo  ''
	@echo  '* Quality:'
	@echo  '- lint            - Phony task that runs all linting tasks'
	@echo  '- test            - Phony task that runs all unit tests'
	@echo  '- fix             - Fixes some linting errors
	@echo  ''
	@echo  '* Release:'
	@echo  '- all (default)   - Default phony task that builds (client and'
	@echo  '                    and server - sided) binaries for development.'
	@echo  '                    Takes an obligatory param `mode` with values'
	@echo  '                    `dev` or `production`.'
	@echo  '- version         - Phony task. Creates changelog from latest'
	@echo  '                    git tag till the latest commit. Creates commit'
	@echo  '                    with given version (as argument) and tags'
	@echo  '                    this commit with this version. Version has to'
	@echo  '                    be passed as `V` argument with ex. `v0.0.0`'
	@echo  '                    format'
	@echo  ''


