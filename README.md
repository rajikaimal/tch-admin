# tch-admin

Create .env for configs

```
mysql="<username>:<password>@tcp(mysql:3306)/tchadmin?charset=utf8mb4&parseTime=True&loc=Local"
```

## Install

### Docker

```
$ docker-compose up --build
```

### Without Docker

```
$ go mod tidy
$ go run migrate/migrate.go
$ go run seed/seed.go
$ go run main.go
```
