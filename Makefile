compose-setup: compose-build
compose:
	docker-compose up

compose-build:
	docker-compose build

compose-lint:
	docker-compose run exercises make lint

docker-local-build:
	docker build --tag hexletbasics/exercises-go .

SUBDIRS := $(wildcard modules/**/*/.)

lint:
	yamllint modules

test: $(SUBDIRS)
$(SUBDIRS):
	@echo
	make test -s -C $@
	@echo

.PHONY: all test $(SUBDIRS)
