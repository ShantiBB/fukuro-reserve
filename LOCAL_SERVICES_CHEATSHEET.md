# Local Services Cheat Sheet

Эта шпаргалка нужна, чтобы быстро и предсказуемо поднимать весь локальный стенд проекта.

## Что поднимаем

- PostgreSQL в Docker
- Базы: `auth`, `hotel`, `booking`
- Миграции для каждого сервиса
- gRPC-сервисы:
  - `auth` на `localhost:8081`
  - `hotel` на `localhost:8082`
  - `booking` на `localhost:8083`
- HTTP gateway:
  - `gateway` на `localhost:8080`

## Важный нюанс

В корневом `.env` лежит:

```env
CONFIG_PATH=./config/local.yaml
```

Из-за этого запуск из корня репозитория работает не для всех сервисов:

- `auth` можно запускать из корня с явным `CONFIG_PATH`
- `hotel` и `booking` удобнее запускать из каталогов самих сервисов
- `gateway` из корня лучше запускать с явным `CONFIG_PATH`

## 1. Поднять PostgreSQL

Из корня репозитория:

```bash
docker compose --env-file .env.dev -f docker/postgres.yaml up -d
```

Проверить, что контейнер жив:

```bash
docker ps --filter name=postgres
```

## 2. Создать базы данных

Если базы ещё не созданы:

```bash
docker exec postgres psql -U postgres -d postgres -c 'CREATE DATABASE "auth";'
docker exec postgres psql -U postgres -d postgres -c 'CREATE DATABASE "hotel";'
docker exec postgres psql -U postgres -d postgres -c 'CREATE DATABASE "booking";'
```

Если база уже существует, Postgres просто сообщит об этом. Альтернатива: сначала проверить список баз:

```bash
docker exec postgres psql -U postgres -d postgres -c '\l'
```

## 3. Накатить миграции

Из корня репозитория:

```bash
task goose service=auth mode=up
task goose service=hotel mode=up
task goose service=booking mode=up
```

Проверка статуса:

```bash
task goose service=auth mode=status
task goose service=hotel mode=status
task goose service=booking mode=status
```

## 4. Запустить сервисы

Рекомендуемый порядок:

1. `auth`
2. `hotel`
3. `booking`
4. `gateway`

### Auth

Из корня репозитория:

```bash
CONFIG_PATH=services/auth/config/local.yaml go run services/auth/cmd/app/main.go
```

Ожидаемый адрес:

```text
localhost:8081
```

### Hotel

Из каталога сервиса:

```bash
cd services/hotel
go run cmd/app/main.go
```

Ожидаемый адрес:

```text
localhost:8082
```

### Booking

Из каталога сервиса:

```bash
cd services/booking
go run cmd/app/main.go
```

Ожидаемый адрес:

```text
localhost:8083
```

### Gateway

Из корня репозитория:

```bash
CONFIG_PATH=services/gateway/config/local.yaml go run services/gateway/cmd/app/main.go
```

Ожидаемый адрес:

```text
localhost:8080
```

## 5. Быстрая проверка

Healthcheck gateway:

```bash
curl -i http://localhost:8080/health
```

Ожидаемый ответ:

```text
HTTP/1.1 200 OK
...
OK
```

## Полная последовательность команд

Если поднимать всё с нуля, рабочая последовательность такая:

```bash
docker compose --env-file .env.dev -f docker/postgres.yaml up -d

docker exec postgres psql -U postgres -d postgres -c 'CREATE DATABASE "auth";'
docker exec postgres psql -U postgres -d postgres -c 'CREATE DATABASE "hotel";'
docker exec postgres psql -U postgres -d postgres -c 'CREATE DATABASE "booking";'

task goose service=auth mode=up
task goose service=hotel mode=up
task goose service=booking mode=up

CONFIG_PATH=services/auth/config/local.yaml go run services/auth/cmd/app/main.go
```

В отдельном терминале:

```bash
cd services/hotel
go run cmd/app/main.go
```

В отдельном терминале:

```bash
cd services/booking
go run cmd/app/main.go
```

В отдельном терминале:

```bash
CONFIG_PATH=services/gateway/config/local.yaml go run services/gateway/cmd/app/main.go
```

## Остановка

Остановить сервисы:

- `Ctrl+C` в каждом терминале

Остановить Postgres:

```bash
docker compose --env-file .env.dev -f docker/postgres.yaml down
```

## Типовые проблемы

### `open ./config/local.yaml: no such file or directory`

Причина:

- сервис стартует не из того каталога
- или ему подставился неправильный `CONFIG_PATH`

Что делать:

- для `hotel` и `booking` запускать из директории сервиса
- для `auth` и `gateway` явно задавать `CONFIG_PATH`

### `bind: address already in use`

Причина:

- сервис уже запущен на нужном порту

Что делать:

- остановить старый процесс
- или проверить, кто слушает порт:

```bash
lsof -i :8080
lsof -i :8081
lsof -i :8082
lsof -i :8083
```

### Ошибки подключения к Postgres

Проверить:

- запущен ли контейнер `postgres`
- созданы ли базы `auth`, `hotel`, `booking`
- накатились ли миграции
- не занят ли порт `5432`

## Короткий prompt для агента

```text
Подними локальный стенд fukuro-reserve: запусти PostgreSQL через docker/postgres.yaml, проверь наличие баз auth/hotel/booking, накатай миграции через task goose, затем запусти auth, hotel, booking и gateway с корректными CONFIG_PATH/working directory. После запуска проверь GET http://localhost:8080/health и сообщи статус каждого сервиса.
```
