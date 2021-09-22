# bankrupt

## docker

```
$ docker pull postgres:12-alpine
$ docker run --name postgres12 -p54332:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
$ docker exec -it postgres12 psql -U root
```
