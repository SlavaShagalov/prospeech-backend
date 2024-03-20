EASYJSON_PATHS = ./internal/...

# ===== RUN =====
.PHONY: stop
stop:
	docker compose stop

.PHONY: down
down:
	docker compose down -v

.PHONY: build
build:
	docker compose -f docker-compose.yml build api

.PHONY: dev-up
dev-up:
	docker compose -f docker-compose.yml up -d --build db redis api nginx

.PHONY: dev-stop
dev-stop:
	docker compose -f docker-compose.yml stop db redis api nginx

.PHONY: prod-up
prod-up:
	cp -r ../prospeech-frontend ./frontend
	docker compose -f docker-compose-prod.yml up -d --build db redis api nginx
	rm -rf frontend

.PHONY: prod-stop
prod-stop:
	docker compose -f docker-compose-prod.yml stop db redis api nginx

# ===== LOGS =====

service = db
.PHONY: logs
logs:
	docker compose logs -f "$(service)"

name = main
.PHONY: logs-api
logs-api:
	tail -f -n +1 "cmd/api/logs/$(name).log" | batcat --paging=never --language=log

# ===== GENERATORS =====

.PHONY: mocks
mocks:
	./scripts/gen_mocks.sh

.PHONY: easyjson
easyjson:
	go generate ${EASYJSON_PATHS}

.PHONY: swag
swag:
	swag init -g cmd/api/main.go

# ===== OTHER =====

.PHONY: format
format:
	swag fmt
