-include /opt/basics/common/common.mk

update-deps:
	go get -u all && go mod tidy

code-lint:
	golangci-lint run modules/...

code-format:
	gofmt -s -w modules

compose-setup: compose-build

compose:
	docker compose up

compose-build:
	docker compose build

compose-bash:
	docker compose run --rm exercises bash

compose-test:
	docker compose run --rm exercises make test

compose-down:
	docker compose down -v --remove-orphans

compose-build-test: compose-build compose

compose-description-lint:
	docker compose run --rm exercises make description-lint

compose-schema-validate:
	docker compose run --rm exercises make schema-validate

ci-check:
	docker compose --file docker-compose.yml build
	docker compose --file docker-compose.yml up --abort-on-container-exit

compose-code-lint:
	docker compose run --rm exercises make code-lint

docker-build:
	docker build -t ghcr.io/hexlet-basics/exercises-go:latest .

docker-push:
	docker push ghcr.io/hexlet-basics/exercises-go:latest
