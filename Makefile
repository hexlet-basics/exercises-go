-include /opt/basics/common/common.mk

code-lint:
	golint -set_exit_status modules/...

compose-setup: compose-build

compose:
	docker compose up

compose-build:
	docker compose build

compose-go-mod:
	docker compose run exercises go mod init exercises || exit 0
	docker compose run exercises go mod vendor

compose-bash:
	docker compose run exercises bash

compose-test:
	docker compose run exercises make test
