DB_URL=mysql://root:123456@tcp(127.0.0.1:3306)/go_fiber

migrate.up:
	migrate -path db/migrate -database "$(DB_URL)" up 1

migrate.down:
	migrate -path db/migrate -database "$(DB_URL)" down 1

gen.model:
	sqlc generate

gen.doc:
	swag init

dev:
	APP_ENV=local air

dev.start:
	APP_ENV=development air

dev.pro:
	APP_ENV=production air

build:
	mkdir build && cp .env.production ./build/.env.production && cp .env.development ./build/.env.development && go build -o ./build/main main.go
