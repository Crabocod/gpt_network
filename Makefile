DB_USER=root
DB_PASSWORD=root
DB_NAME=gpt_network
DB_HOST=localhost
DB_PORT=5432

up:
	docker-compose up -d

down:
	docker-compose down

migrate-up:
	migrate -path internal/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path internal/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down