# bankrupt

## docker

```
$ docker pull postgres:12-alpine
$ docker run --name postgres12 -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
$ docker exec -it postgres12 psql -U root
```

```
$ brew install golang-migrate
$ brew install sqlc
```

```
$ make postgres
$ make createdb
$ make dropdb
$ make migrateup
$ make migratedown
```

```
$ go get github.com/lib/pq
$ go get github.com/stretchr/testify
$ go mod tidy
$ go get -u github.com/gin-gonic/gin
$ go get github.com/spf13/viper
```
