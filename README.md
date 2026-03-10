# Echo Request Logger

HTTP echo-сервис на Go, который логирует информацию о каждом входящем запросе в PostgreSQL.

## Требования

### Стек технологий
- **Язык**: Go
- **ORM**: [Ent](https://entgo.io/) — генерация моделей и схем
- **Миграции**: [Atlas](https://atlasgo.io/) — версионные миграции БД
- **БД**: PostgreSQL
- **Контейнеризация**: Docker + Docker Compose
- **CI**: GitHub Actions

### Функциональные требования
- Сервис принимает любые HTTP-запросы и возвращает эхо-ответ (метод, путь, заголовки, тело)
- Каждый запрос логируется в таблицу `request_logs` в PostgreSQL
- Записываются: метод, путь, заголовки, тело запроса, IP-адрес, время запроса

### Инфраструктурные требования
- Проект поднимается одной командой: `make up`
- Docker Compose оркестрирует PostgreSQL и приложение
- Миграции применяются автоматически при старте (через Atlas)
- Файлы миграций хранятся в формате LF (настроено через `.gitattributes`)
- CI pipeline в GitHub Actions: линтинг, тесты, сборка Docker-образа

### Структура проекта
```
.
├── .github/workflows/    # GitHub Actions CI
├── cmd/server/           # Точка входа приложения
├── ent/                  # Ent схемы и сгенерированный код
│   └── schema/           # Определения сущностей
├── internal/handler/     # HTTP обработчики
├── migrations/           # Atlas миграции (LF)
├── atlas.hcl             # Конфигурация Atlas
├── Dockerfile            # Multi-stage сборка
├── docker-compose.yml    # Оркестрация сервисов
├── Makefile              # Команды управления проектом
└── README.md
```

## Быстрый старт

```bash
make up
```

Сервис будет доступен на `http://localhost:8080`.

### Примеры запросов

```bash
# GET запрос
curl http://localhost:8080/hello

# POST запрос с телом
curl -X POST http://localhost:8080/api/data -H "Content-Type: application/json" -d '{"key":"value"}'
```

## Makefile команды

| Команда | Описание |
|---------|----------|
| `make up` | Собрать и запустить все сервисы |
| `make down` | Остановить все сервисы |
| `make build` | Собрать Docker-образы |
| `make logs` | Показать логи сервисов |
| `make migrate-new` | Создать новую миграцию |
| `make generate` | Сгенерировать Ent код |
