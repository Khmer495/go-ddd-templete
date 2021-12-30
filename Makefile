.PHONY: clean-local-mysql
clean-local-mysql:
	rm -rf mysql/data
	mkdir mysql/data

DOCKER_COMPOSE_INFRA_FILE=local.infra.docker-compose.yaml
DOCKER_COMPOSE_INFRA_NETWORK=infrastructure
.PHONY: run-infra-local
run-infra-local:
	@if [ -z "`docker network ls | grep $(DOCKER_COMPOSE_INFRA_NETWORK)`" ]; then docker network create $(DOCKER_COMPOSE_INFRA_NETWORK); fi
	docker compose -f $(DOCKER_COMPOSE_INFRA_FILE) up

.PHONY: run-infra-local-build
run-infra-local-build:
	docker compose -f $(DOCKER_COMPOSE_INFRA_FILE) build --force-rm --no-cache

.PHONY: run-infra-local-down
run-infra-local-down:
	docker compose -f $(DOCKER_COMPOSE_INFRA_FILE) down
