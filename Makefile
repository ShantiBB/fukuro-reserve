mock-gen-auth:
	rm -rf services/auth/internal/mocks
	cd services/auth && mockery

swag-gen-auth:
	swag init -g services/auth/cmd/app/main.go --output services/auth/docs

test-auth-handler:
	go test ./services/auth/internal/http/handler/

run-%:
	CONFIG_PATH=$(shell pwd)/services/$*/config/local.yaml go run ./services/$*/cmd/app/main.go

build-auth:
	go build -o build/auth-service services/auth/cmd/app/main.go

build-hotel:
	go build -o build/hotel-service services/hotel/cmd/app/main.go

postgres-up:
	docker compose --env-file docker/.env -f docker/postgres.yaml up -d
postgres-stop:
	docker compose --env-file docker/.env -f docker/postgres.yaml down
