# Fukuro-reserve
> микросервисная система бронирования отелей, написанная на Go.

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17+-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-enabled-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![Swagger](https://img.shields.io/badge/Swagger-API%20Docs-85EA2D?style=flat&logo=swagger)](https://swagger.io/)
[![GolangCI-Lint](https://img.shields.io/badge/GolangCI--Lint-linter-00ADD8?style=flat&logo=go)](https://github.com/golangci/golangci-lint)
[![Task](https://img.shields.io/badge/Task-runner-29BEB0?style=flat&logo=task)](https://taskfile.dev/)

### Технологии
- Backend - go (chi + pgx)
- Repository - PostgreSQL 17
- Documentation - swag
- Mock generation - mockery
- Migrator lib - golang-migrate

### Расширение
- Добавить сервис для обработки уведомлений (интеграция kafka)
- Добавить сервис для обработки отзывов и комментариев.
- Добавить небольшой frontend с интеграцией карт (OSM, Mapbox, Яндекс API)
