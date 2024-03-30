.PHONY: build-run
build-run:
	@docker compose up --build

.PHONY: run
run:
	@docker compose up

.PHONY: stop
stop:
	@docker compose down