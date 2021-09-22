# bankrupt

## docker

```
$ docker pull postgres:12-alpine
$ docker run --name postgres12 -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=xxxxxx -d postgres:12-alpine
$ docker exec -it postgres12 psql -U root
```
