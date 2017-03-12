DOCKER_REPO ?= rinq/test-server
.DEFAULT_GOAL = run

SHELL := /bin/bash
-include artifacts/make/go.mk

.PHONY:
run: docker
	-docker network create --driver overlay rinq-test
	-docker service rm $(shell docker service ls -q --filter name=rinq-test)

	docker service create \
		--name rinq-test-server \
		--network rinq-test \
		--env RINQ_AMQP_DSN=amqp://rinq-test-rabbitmq \
		--env RINQ_HTTPD_ORIGIN="*" \
		$(DOCKER_REPO):dev

	docker pull rabbitmq:alpine
	docker service create \
		--name rinq-test-rabbitmq \
		--network rinq-test \
		rabbitmq:alpine

	docker pull rinq/httpd:dev
	docker service create \
		--name rinq-test-httpd \
		--network rinq-test \
		--publish 80:8080 \
		--env RINQ_AMQP_DSN=amqp://rinq-test-rabbitmq \
		--env RINQ_HTTPD_ORIGIN="*" \
		rinq/httpd:dev

.PHONY:
docker-clean::
	-docker service rm $(shell docker service ls -q --filter name=rinq-test) 2>/dev/null
	@sleep 1
	-docker network rm rinq-test 2>/dev/null
	-docker image rm "$(DOCKER_REPO):$(DOCKER_TAG)"

artifacts/make/%.mk:
	bash <(curl -s https://rinq.github.io/make/install) $@
