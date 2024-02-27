DB_CONTAINER_NAME = go_webserver-postgres-1
DB_USER = postgres
DB_PASSWORD = postgres
DB_ENV = development
DB_NAME = go_webserver_$(DB_ENV)

db/migrate/new:
	migrate create -ext sql -dir db/migrations/ $(NAME)
db/migrate/up:
	migrate -path db/migrations/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up
	make db/schema
db/migrate/down:
	migrate -path db/migrations/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down
	make db/schema
db/migrate/goto:
	migrate -path db/migrations/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/${DB_NAME}?sslmode=disable" goto $(VERSION)
	make db/schema
db/migrate/fix:
	migrate -path db/migrations/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/${DB_NAME}?sslmode=disable" force $(VERSION)
db/drop:
	docker exec -it $(DB_CONTAINER_NAME) dropdb -U $(DB_USER) $(DB_NAME)
db/create:
	docker exec -it $(DB_CONTAINER_NAME) createdb -U $(DB_USER) $(DB_NAME)
	make db/schema
db/schema:
	docker exec -it $(DB_CONTAINER_NAME) pg_dump -U postgres --dbname=$(DB_NAME) --schema-only --no-owner --no-acl > db/schema.sql
db/setup:
	make db/create
	make db/migrate/up
test/run:
	ENV=test go test ./...
