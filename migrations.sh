#!bin/bash
source .env

export MIGRATION_DSN="host=pg port=5432 dbname=order user=order-user password=secret sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v