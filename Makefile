GOCMD=go
GOLINT=golint
MAKE=make
NPM=npm
mode=prod

.DEFAULT_GOAL := all

.PHONY: install
install:
	$(NPM) install
	$(NPM) run elm:package:install --yes
	$(GOCMD) get github.com/coda-it/gowebserver
	$(GOCMD) get github.com/influxdata/influxdb1-client/v2
	$(GOCMD) get gopkg.in/mgo.v2
	$(GOCMD) get github.com/golang/lint/golint

.PHONY: all
all:
	$(GOCMD) generate
	$(GOCMD) build -o smarthome
	$(NPM) run build:$(mode)

.PHONY: test
test:
	$(NPM) run test:unit
	$(GOCMD) test ./...

.PHONY: lint
lint:
	$(NPM) run lint
	$(NPM) run csslint
	$(GOLINT) ./...
	$(GOCMD) vet ./...

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


