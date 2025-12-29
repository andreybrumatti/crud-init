include .env
export

migrate-up:
	migrate -path internal/database/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" -verbose up

migrate-down:
	migrate -path internal/database/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" -verbose down