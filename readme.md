# WB Order Service

## Описание

Это сервис для обработки заказов, реализованный на Go.  
Он использует PostgreSQL для хранения данных, Kafka для асинхронной обработки, и предоставляет HTTP API для работы с заказами.  
Веб-интерфейс доступен через статическую страницу (`web/index.html`).

## Структура проекта

- `cmd/main.go` — точка входа приложения
- `internal/config` — конфигурация приложения
- `internal/model` — модели данных
- `internal/repository` — работа с базой данных
- `internal/service` — бизнес-логика
- `internal/handlers` — HTTP-обработчики
- `internal/kafka` — работа с Kafka
- `web/` — статические файлы (веб-страница)

## Запуск

1. Запустите PostgreSQL и Kafka (можно через Docker Compose):

   ```sh
   docker compose up -d
   ```

2. Запустите приложение:

   ```sh
   make run
   ```

3. Откройте веб-страницу:
   - [http://localhost:8080/](http://localhost:8080/)

## Конфигурация

- Настройки хранятся в файле `.env`
- Примеры переменных:
  ```
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=postgres
  DB_PASSWORD=yourpassword
  DB_NAME=wb
  HTTP_PORT=8080
  KAFKA_BROKERS=localhost:9092
  KAFKA_TOPIC=orders
  KAFKA_GROUP_ID=order-group
  ```

## API

- `GET /order/{uid}` — получить заказ по UID
- `POST /order` — создать заказ

## Веб-интерфейс

- Статическая страница: [`web/index.html`](web/index.html)
