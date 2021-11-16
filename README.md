# bankrupt

## docker

```
$ docker pull postgres:12-alpine
$ docker run --name postgres12 -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
$ docker exec -it postgres12 psql -U root
$ docker build -t bankrupt:latest .
$ docker images
```

## mac

```
$ brew install golang-migrate
$ brew install sqlc
```

# setup

```
$ make postgres
$ make createdb
$ make dropdb
$ make migrateup
$ make migratedown
```

# go

```
$ go get github.com/lib/pq
$ go get github.com/stretchr/testify
$ go mod tidy
$ go get -u github.com/gin-gonic/gin
$ go get github.com/spf13/viper
$ go install github.com/golang/mock/mockgen@v1.6.0
$ go get github.com/google/uuid
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/o1egl/paseto
```

# vscode

```
"go.testFlags": ["-v"],
```
