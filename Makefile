gen-migrate: #name=<migration_name>
	migrate create -ext sql -dir ./services/auth/migrations -seq $(name)

add-migrate: # service=<service_name> name=<migration_name>
	migrate create -ext sql -dir ./services/$(service)/migrations -seq $(name)

run-%:
	CONFIG_PATH=$(shell pwd)/services/$*/config/local.yaml go run ./services/$*/cmd/app/main.go

build-%:
	go build -o build/$*-service services/$*/cmd/app/main.go

mock-gen-%:
	rm -rf services/$*/internal/mocks
	cd services/$* && mockery

swag-gen-%:
	cd services/$* && swag init --parseDependency -g cmd/app/main.go --output docs

test-unit-auth-handler:
	go test ./services/auth/internal/http/handler/

postgres-up:
	docker compose --env-file docker/.env -f docker/postgres.yaml up -d
postgres-stop:
	docker compose --env-file docker/.env -f docker/postgres.yaml down
