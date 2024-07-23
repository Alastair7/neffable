#!/bin/bash
set -e

# Wait for PostgreSQL to be ready
until pg_isready -U "$POSTGRES_USER"; do
  echo "Waiting for database to be ready..."
  sleep 2
done

# Check if the database exists
if psql -U "$POSTGRES_USER" -lqt | cut -d \| -f 1 | grep -qw "$POSTGRES_DB"; then
    echo "Database $POSTGRES_DB already exists."
else
    echo "Creating database $POSTGRES_DB..."
    createdb -U "$POSTGRES_USER" "$POSTGRES_DB"
fi