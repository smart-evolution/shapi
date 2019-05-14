GOCMD=go
GOLINT=golint
GOFMT=gofmt
MAKE=make
NPM=npm
mode=prod

.DEFAULT_GOAL := all

.PHONY: install
install:
	$(NPM) install
	$(GOCMD) get github.com/coda-it/gowebserver
	$(GOCMD) get github.com/influxdata/influxdb1-client/v2
	$(GOCMD) get gopkg.in/mgo.v2
	$(GOCMD) get golang.org/x/lint/golint

.PHONY: all
all:
	$(GOCMD) generate
	$(GOCMD) build -o smarthome
	$(NPM) run build:$(mode)

.PHONY: test
test:
	$(NPM) run test
	$(GOCMD) test ./...

.PHONY: lint
lint:
	$(NPM) run flow
	$(NPM) run lint
	$(NPM) run csslint
	./scripts/gofmt_test.sh
	$(GOLINT) ./...
	$(GOCMD) vet ./...

.PHONY: fix
fix:
	$(NPM) run prettify
	$(NPM) run lint:fix
	$(NPM) run csslint:fix
	$(GOFMT) -w .

.PHONY: version
version:
	git tag $(V)
	./scripts/changelog.sh
	go generate
	$(NPM) version $(V) --no-git-tag-version
	git add package.json
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


