chmod +x bin/gen-file-folder.sh

# Create Folder

mkdir cmd

mkdir cmd/http
touch cmd/http/main.go

mkdir docs

mkdir infrastructures
mkdir infrastructures/db
mkdir infrastructures/db/gorm
touch infrastructures/config.go
touch infrastructures/db_migration.go
touch infrastructures/entity_registry.go

mkdir internal

mkdir internal/adapters
mkdir internal/adapters/cache
mkdir internal/adapters/cache/redis
touch internal/adapters/cache/redis/redis.go

mkdir internal/adapters/handler
mkdir internal/adapters/handler/controllers
mkdir internal/adapters/handler/routers

mkdir internal/adapters/token

mkdir internal/core
mkdir internal/core/domain
mkdir internal/core/ports
mkdir internal/core/services
mkdir internal/core/utils
