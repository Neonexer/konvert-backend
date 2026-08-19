include .env
export

export PROJECT_ROOT=$(shell pwd)


env-up:
	docker compose up -d konvert-postgres

env-down:
	docker compose down konvert-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
		if [ "$$ans" = "y" ]; then \
			docker compose down konvert-postgres && \
			rm -rf out/pgdata && \
			echo "Файлы окружения очищены"; \
		else \
			echo "Очистка окружения отменена"; \
		fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm konvert-postgres-migrate \
		create \
		-ext sql \
		-dir migrations \
		-seq "$(seq)"

