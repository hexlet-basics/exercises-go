-include /opt/basics/common/common.mk

compose-setup: compose-build

compose:
	docker-compose up

compose-build:
	docker-compose build

code-lint:
	golint -set_exit_status modules/...

compose-bash:
	docker-compose run exercises bash

compose-test:
	docker-compose run exercises make test
