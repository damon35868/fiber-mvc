DB_URL=mysql://root:3586878@tcp(127.0.0.1:3306)/go_fiber

migrateup:
	migrate -path db/migrate -database "$(DB_URL)" up 1

migratedown:
	migrate -path db/migrate -database "$(DB_URL)" down 1

modelgen:
	sqlc generate

docgen:
	swag init

dev:
	APP_ENV=local air

devstart:
	APP_ENV=development air

devpro:
	APP_ENV=production air

build:
	mkdir build && cp .env.production ./build/.env.production && cp .env.development ./build/.env.development && go build -o ./build/main main.go