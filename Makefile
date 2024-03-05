EASYJSON_PATHS = ./internal/...

# ===== RUN =====
.PHONY: up
up:
	make api-up

.PHONY: stop
stop:
	make api-stop

.PHONY: down
down:
	docker compose down -v

.PHONY: api-up
api-up:
	docker compose -f docker-compose.yml up -d --build db redis api nginx

.PHONY: api-stop
api-stop:
	docker compose -f docker-compose.yml stop db redis api nginx

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

# ===== FORMAT =====

.PHONY: format
format:
	swag fmt
