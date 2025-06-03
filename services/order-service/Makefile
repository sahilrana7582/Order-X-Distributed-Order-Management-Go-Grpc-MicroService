DB_URL=postgres://postgres:1234@localhost:5432/orderX?sslmode=disable
MIGRATIONS_DIR=./internal/db/migrations



# migrate create -ext sql -dir ./internal/db/migrations -seq


# migrate create -ext sql -dir ./internal/db/migrations -seq order_up


run:
	go run ./cmd/order-service/main.go


migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force 1