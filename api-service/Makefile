DB_USER=root
DB_PASSWORD=root
DB_NAME=gpt_network
DB_HOST=localhost
DB_PORT=5432


migrate-up:
	docker-compose exec db bash -c '\
	for file in /docker-entrypoint-initdb.d/migrations/*.up.sql; do \
		echo "Running migration: $$file"; \
		psql -U $(DB_USER) -d $(DB_NAME) -f "$$file"; \
	done'

migrate-down:
	docker-compose exec db bash -c '\
	for file in /docker-entrypoint-initdb.d/migrations/*.down.sql; do \
		echo "Rolling back migration: $$file"; \
		psql -U $(DB_USER) -d $(DB_NAME) -f "$$file"; \
	done'