#!/bin/bash

set -e

# 環境変数の設定
DATABASE_URL="postgres://${POSTGRES_USER:-postgres}:${POSTGRES_PASSWORD:-password}@postgres:${POSTGRES_PORT:-5432}/${POSTGRES_DB:-emopathy}?sslmode=disable"

echo "Database URL: $DATABASE_URL"

echo "Waiting for PostgreSQL to be ready..."
until pg_isready -h postgres -p 5432 -U postgres; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done


echo "Running migrations..."
echo "Migration files:"
ls -la /migrations/

migrate -path /migrations -database "$DATABASE_URL" up

echo "Migrations completed successfully!"
