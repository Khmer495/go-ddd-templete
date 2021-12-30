.PHONY: generate
install-tools:
	./script/install_tools.sh

generate:
	./script/ent/generate.sh
	./script/mockgen/generate.sh
	./script/oapi-codegen/v1/generate.sh

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./... -v

.PHONY: run-api-local
run-api-local:
	docker compose -f local.api.docker-compose.yaml up

.PHONY: run-api-local-build
run-api-local-build:
	docker compose -f local.api.docker-compose.yaml build --force-rm --no-cache

.PHONY: run-api-local-down
run-api-local-down:
	docker compose -f local.api.docker-compose.yaml down
