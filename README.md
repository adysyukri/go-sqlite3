# Playing around with Sqlite3

## Have these installed first
```
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

```

## Migrate command
```
migrate create -ext sql -dir db/migrations -seq create_table

migrate -database sqlite3://test.db -path db/migrations up


migrate -database sqlite3://test.db -path db/migrations down

```