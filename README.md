# Golang + GraphQL API

This is a fun project that will be used on the ValhallaCoders Blog.

## Dev Dependencies

- Golang [Dep](https://golang.github.io/dep/docs/installation.html): packages/dependencies manager. To add more dependencies: `$ dep ensure -add github.com/foo/bar github.com/bar/foo`
- Postgres DB Administrtor [pgAdmin](https://www.pgadmin.org/download/)

## Dev

<!-- The DB and the API must run in the same network, to create the isolated network, run:`$ docker network create isolated` -->

### DB

- Just run `$ sh db-start.sh` to start the Alpine-postgres DB container.

### API

- After a `dep ensure`, just run `$ go run main.go`... for now
- If you need to add a new golang dependency, just import it and run `$ dep ensure`. Warning: if you run this command will install all dependencies that are required and delete the ones thar are installed but not in use.

## Moving to Prod

- [] automatize this process.

```go
fmt.Println("See you in Valhalla!")
```
