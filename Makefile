.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down --rmi all --volumes --remove-orphans

.PHONY: restart
restart:
	docker compose restart

.PHONY: tidy
tidy:
	docker compose exec app go mod tidy

.PHONY: lint
lint:
	docker compose exec app golangci-lint run

.PHONY: format
format:
	docker compose exec app golangci-lint run --fix

.PHONY: test
test:
	docker compose exec app go test -v ./...