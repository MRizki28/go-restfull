run:
	go run ./cmd

dev:
	go run ./cmd

build:
	go build -o app ./cmd


DB_URL=mysql://root:secret@tcp(127.0.0.1:3306)/go_fiber

migration-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir database/migrations $$name

migration-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migration-down:
	migrate -path database/migrations -database "$(DB_URL)" down 1
