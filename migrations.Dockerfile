FROM lynqtech/docker-sql-migrate:1.0.5

COPY migrations /migrations

WORKDIR /migrations
ENTRYPOINT ["sql-migrate", "up", "-env=deployment"]
