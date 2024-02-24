# go_webserver

## Installation
```bash
go mod tidy # install dependencies
brew install golang-migrate # install migration tool
docker-compose up -d
make db/create
make db/migrate/up
```
