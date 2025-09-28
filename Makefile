run:
	go run ./main.go

postgres-run:
	docker compose --env-file .env -f docker/postgres-compose.yaml up -d
postgres-stop:
	docker compose --env-file .env -f docker/postgres-compose.yaml down
