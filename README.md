# bankrupt

## docker

```
$ docker pull postgres:12-alpine
$ docker run --name postgres12 -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
$ docker exec -it postgres12 psql -U root
$ docker build -t bankrupt:latest .
$ docker images
$ docker rmi [IMAGE ID]
$ docker run --name bankrupt -p 8888:8888 bankrupt:latest
$ docker run --name bankrupt -p 8888:8888 -e GIN_MODE=release bankrupt:latest
$ docker rm bankrupt
$ docker container inspect postgres12
$ docker container inspect bankrupt
$ docker network ls
$ docker network inspect bridge
$ docker network create bank-network
$ docker network connect bank-network postgres12
$ docker run --name bankrupt --network bank-network -p 8888:8888 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.18.0.2:5432/simple_bank?sslmode=disable" bankrupt:latest
$ docker run --name bankrupt --network bank-network -p 8888:8888 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" bankrupt:latest
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

# aws

```
$ aws configure
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text | jq 'to_entries'
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text | jq 'to_entries|map(.key)'
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text | jq 'to_entries|map(.value)'
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text | jq 'to_entries|map("\(.key)=\(.value)")'
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text | jq 'to_entries|map("\(.key)=\(.value)")|.[]'
$ aws secretsmanager get-secret-value --secret-id bankrupt --profile github-ci --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]'
$ aws ecr get-login-password --profile github-ci | docker login --username AWS --password-stdin [arn]
```

# vscode

```
"go.testFlags": ["-v"],
```

# misc

```
$ openssl rand -hex 64 | head -c 32
```
