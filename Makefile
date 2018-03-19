GOCMD=go
GOGET=$(GOCMD) get
GOGENERATE=$(GOCMD) generate
GOBUILD=$(GOCMD) build -o smarthome
GOTEST=$(GOCMD) test ./...

NPM=npm
NPMINSTALL=$(NPM) install
NPMBUILD=$(NPM) run build

ELMPKGINSTALL=npm run elm:package:install --yes

.DEFAULT_GOAL := all

.PHONY: install
install:
	$(NPMINSTALL)
	$(ELMPKGINSTALL)
	$(GOGET) github.com/oskarszura/gowebserver

.PHONY: all
all:
	$(GOGENERATE)
	$(GOBUILD)
	$(NPMBUILD)

.PHONY: test
test:
	$(GOTEST)

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
	@echo  'Available tasks:'
	@echo  '* Installation:'
	@echo  '- install         - Phony task that installs all required (client'
	@echo  '                    and server - sided) dependencies'
	@echo  ''
	@echo  '* Build:'
	@echo  '- all (default)   - Default phony task that builds (client and'
	@echo  '                    and server - sided) binaries'
	@echo  ''
	@echo  '* Tests:'
	@echo  '- test (default)   - Phony task that runs all unit tests'
	@echo  ''
	@echo  '* Release:'
	@echo  '- version         - Phony task. Creates changelog from latest'
	@echo  '                    git tag till the latest commit. Creates commit'
	@echo  '                    with given version (as argument) and tags'
	@echo  '                    this commit with this version. Version has to'
	@echo  '                    be passed as `V` argument with ex. `v0.0.0`'
	@echo  '                    format'
	@echo  ''


