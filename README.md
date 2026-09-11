# Бэкенд для сервиса "Конверт"

## Документация

- [Сущности](./docs/entities.md)
- [Ручки](./docs/handlers.md)

## Деплой

При пуше в `main` GitHub Actions запускает `go vet`, сборку и тесты. Если
проверки успешны, workflow подключается к VPS по SSH, обновляет
checkout до `origin/main`, генерирует Swagger и запускает `make konvert-deploy`.

На VPS должны быть установлены Git и Docker с Compose plugin; Go не требуется.
Один раз клонируйте репозиторий в каталог, который будет указан в `VPS_APP_DIR`,
и настройте для него SSH-ключ или другой способ доступа к GitHub, чтобы команда
`git fetch origin main` могла выполниться без вопросов. В этом каталоге заранее
создайте локальный `.env` из `.env.example` и убедитесь, что пользователь SSH
может выполнять Docker-команды.

Добавьте в настройках репозитория `Settings -> Secrets and variables -> Actions`
следующие Repository variables:

- `VPS_HOST` - адрес VPS, например `203.0.113.10`;
- `VPS_PORT` - SSH-порт, необязательно, по умолчанию `22`;
- `VPS_USER` - пользователь SSH;
- `VPS_APP_DIR` - абсолютный путь к checkout репозитория на VPS.

В Repository secrets добавьте только `VPS_PASSWORD`.
