# Go

## Pre-requiesites

1. Instal Docker Compose
2. Run 
```bash
docker compose up
```
or
```bash
docker-compose up
```
if you don't see changes that you made, add 
```bash
--build
```

## DB connection

Use `https://gorm.io/docs/` to configure connection to the database.
Once you start the container, you will be able to connect to the database at `postgresql://postgres:postgres@localhost:5432/riskengine?sslmode=disable` and database name: `postgres`