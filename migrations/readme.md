# Migrations

In the migrations folder we place all our db migration files.
In order to run a migration on local machine you can run following commands:

- to install `sql-migrate`

```
go get -v github.com/rubenv/sql-migrate/...
```

- to apply migrations

```
sql-migrate up -env="local"
```
