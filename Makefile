-include /opt/basics/common/common.mk

code-lint:
	golint -set_exit_status modules/...

code-format:
	gofmt -s -w modules

compose-setup: compose-build

compose:
	docker compose up

compose-build:
	docker compose build

compose-bash:
	docker compose run exercises bash

compose-test:
	docker compose run exercises make test

compose-build-test: compose-build compose
