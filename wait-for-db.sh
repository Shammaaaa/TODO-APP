#!/bin/sh

if [ "$#" -lt 2 ]; then
    echo "Usage: $0 <host> <port> [command] [args...]"
    exit 1
fi

host="$1"
port="$2"
shift 2

if [ -z "$POSTGRES_USER" ] || [ -z "$POSTGRES_DB" ] || [ -z "$POSTGRES_PASSWORD" ]; then
    echo "Error: POSTGRES_USER, POSTGRES_DB or POSTGRES_PASSWORD is not set."
    exit 1
fi

until PGPASSWORD=$POSTGRES_PASSWORD psql --quiet --no-align --tuples-only --single-transaction \
       --host "$host" \
       --port "$port" \
       --username "$POSTGRES_USER" \
       --dbname "$POSTGRES_DB" \
       --command '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up"

if [ "$#" -gt 0 ]; then
    >&2 echo "Executing command: $*"
    if ! eval "$@"; then
        >&2 echo "Failed to execute command."
        exit $?
    fi
fi