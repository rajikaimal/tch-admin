# tch-admin

Create .env at the root of the project for config

Example:

```
mysql="<username>:<password>@tcp(mysql:3306)/tchadmin?charset=utf8mb4&parseTime=True&loc=Local"
```

## Install

### Docker

```
$ docker-compose up --build
```

Import `init.sql` to mysql

### Without Docker

```
$ go mod tidy
$ go run migrate/migrate.go
$ go run seed/seed.go
$ go run main.go
```
