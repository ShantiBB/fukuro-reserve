run-auth:
	CONFIG_PATH=$(shell pwd)/services/auth/config/local.yaml go run ./services/auth/cmd/app/main.go
run-hotel:
	CONFIG_PATH=$(shell pwd)/services/hotel/config/local.yaml go run ./services/hotel/cmd/app/main.go

postgres-run:
	docker compose --env-file docker/.env -f docker/postgres.yaml up -d
postgres-stop:
	docker compose --env-file docker/.env -f docker/postgres.yaml down
