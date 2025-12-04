auth-run:
	go run ./services/auth/cmd/auth/main.go

postgres-run:
	docker compose --env-file .env -f docker/postgres-compose.yaml up -d
postgres-stop:
	docker compose --env-file .env -f docker/postgres-compose.yaml down
