# Gateway Smoke Test

Запуск из корня репозитория:

```bash
bash scripts/gateway_smoke.sh
```

Что делает скрипт:

- проверяет `/health`
- создаёт временного `owner` и временного `admin`
- повышает временного `admin` напрямую в БД `auth`
- прогоняет все ручки `gateway`:
  - `auth`
  - `users`
  - `hotels`
  - `rooms`
  - `bookings`
- удаляет за собой временные сущности

Требования:

- запущены `auth`, `hotel`, `booking`, `gateway`
- поднят контейнер Postgres `postgres`
- доступны `jq`, `python3`, `docker`
