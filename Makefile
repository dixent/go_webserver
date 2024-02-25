DB_CONTAINER_NAME = go_webserver-postgres-1
DB_USER = postgres
DB_NAME = go_webserver

db/migrate/new:
	migrate create -ext sql -dir db/migrations/ -seq $(NAME)
db/migrate/up: db/schema
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/go_webserver?sslmode=disable" -verbose up
db/migrate/down: db/schema
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/go_webserver?sslmode=disable" -verbose down
db/migrate/goto: db/schema
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/go_webserver?sslmode=disable" goto $(VERSION)
db/migrate/fix:
	migrate -path db/migrations/ -database "postgresql://postgres:postgres@localhost:5432/go_webserver?sslmode=disable" force $(VERSION)
db/drop:
	docker exec -it $(DB_CONTAINER_NAME) dropdb -U $(DB_USER) $(DB_NAME)
db/create: db/schema
	docker exec -it $(DB_CONTAINER_NAME) createdb -U $(DB_USER) $(DB_NAME)
db/schema:
	docker exec -it go_webserver-postgres-1 pg_dump -U postgres --dbname=go_webserver --schema-only --no-owner --no-acl > db/schema.sql
