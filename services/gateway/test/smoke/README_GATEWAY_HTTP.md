# Gateway HTTP Smoke

HTTP smoke-набор для `gateway` находится в:
- `services/gateway/test/smoke/cases/`

Файлы сценариев:
- `00_setup.http`
- `01_auth.http`
- `02_hotels.http`
- `03_rooms.http`
- `04_bookings.http`
- `05_cleanup.http`

## Переменные окружения

Используется один файл:
- `services/gateway/test/smoke/http-client.env.json`

`adminEmail/adminPassword` в нем тестовые (dev-only), отдельный private env-файл не нужен.

## Запуск

Из корня репозитория:
```bash
task gateway-smoke
```

С выбором environment (по умолчанию `local`):
```bash
task gateway-smoke env=local
```

Эквивалентный запуск через `ijhttp`:
```bash
ijhttp \
  services/gateway/test/smoke/cases/*.http \
  -e local \
  -v services/gateway/test/smoke/http-client.env.json \
  -L BASIC
```

## Что проверяет smoke

- auth/users flow
- hotels CRUD flow
- rooms CRUD flow
- bookings flow
- cleanup созданных сущностей

## Важно

- `00_setup.http` генерирует runtime-переменные (token/id), поэтому сценарии нужно запускать только полным набором или строго по порядку.
- Если запускать отдельный файл из середины, будут ошибки по отсутствующим переменным.
