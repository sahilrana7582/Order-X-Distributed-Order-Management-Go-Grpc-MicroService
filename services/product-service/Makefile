DB_URL=postgres://postgres:1234@localhost:5432/orderX?sslmode=disable
MIGRATIONS_DIR=./internal/db/migrations

run:
	go run ./cmd/product-service/main.go



migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force 1