compose:
	docker-compose up

compose-build:
	docker-compose build

SUBDIRS := $(wildcard modules/**/*/.)

lint:
	yamllint modules

test: $(SUBDIRS)
$(SUBDIRS):
	@echo
	make test -s -C $@
	@echo

.PHONY: all test $(SUBDIRS)
