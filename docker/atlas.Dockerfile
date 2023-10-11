FROM arigaio/atlas:latest

WORKDIR /migrations

COPY ent/migrate/migrations .
