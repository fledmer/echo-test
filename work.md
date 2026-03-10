# Work Plan

## Этап 1: Инициализация проекта
- [x] Создать Go module (`go mod init`)
- [x] Создать README.md с требованиями
- [x] Создать work.md (этот файл)

## Этап 2: Модели данных (Ent)
- [x] Установить Ent
- [x] Создать схему `RequestLog` с полями:
  - `method` (string) — HTTP метод
  - `path` (string) — путь запроса
  - `headers` (string/JSON) — заголовки
  - `body` (string) — тело запроса
  - `ip` (string) — IP адрес клиента
  - `created_at` (time) — время запроса
- [x] Сгенерировать Ent код

## Этап 3: HTTP сервис
- [x] Создать `cmd/server/main.go` — точка входа
- [x] Создать `internal/handler/echo.go` — обработчик эхо-запросов
- [x] Подключить Ent клиент к PostgreSQL
- [x] Реализовать логирование запросов в БД

## Этап 4: Миграции (Atlas)
- [x] Создать `atlas.hcl` конфигурацию
- [x] Сгенерировать начальную миграцию
- [x] Настроить `.gitattributes` для LF в файлах миграций

## Этап 5: Docker
- [x] Написать `Dockerfile` (multi-stage build)
- [x] Написать `docker-compose.yml` (postgres + app + atlas migrate)
- [x] Проверить, что `make up` поднимает всё

## Этап 6: CI (GitHub Actions)
- [x] Workflow: lint (golangci-lint)
- [x] Workflow: test (`go test ./...`)
- [x] Workflow: build Docker image
- [x] Workflow: проверка миграций

## Этап 7: Финализация
- [x] Проверить структуру проекта
- [x] Проверить, что всё собирается
- [x] Убедиться, что миграции в LF формате
