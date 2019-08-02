# Golang + GraphQL API

This is a fun project that will be used on the ValhallaCoders Blog.

## Dev Dependencies

- Postgres DB Administrtor [pgAdmin](https://www.pgadmin.org/download/)
- [Dep](https://golang.github.io/dep/docs/installation.html): Golang packages/dependencies manager. To add more dependencies: `$ dep ensure -add github.com/foo/bar github.com/bar/foo`. If you need to add a new golang dependency, just import it and run `$ dep ensure`. **Warning**: if you run this command will install all dependencies that are required and delete the ones thar are installed but not in use.
- [Realize](https://github.com/oxequa/realize): Golang configurable live reload and task runner. `$ go get https://github.com/oxequa/realize` to grab the binary. PS: since this is a program, it cannot be on dep dependencies file.

## Dev

<!-- The DB and the API must run in the same network, to create the isolated network, run:`$ docker network create isolated` -->

### DB

- Just run `$ sh db-start.sh` to start the Alpine-postgres DB container.

### API

- After a `dep ensure`, just run `$ realize start`, this will execute the `main.go` and watch for changes to hotreload them.

## TODO

- ~~add a dependency manager~~
- ~~add a task runner/hotreload to watch golang code~~
- ~~make the API reads env vars~~
- add some response to /
- clean the code
- release the v0.0
- release the v0.1 (containerized)
- implement http2

## Moving to Prod

- automatize this process.

```go
fmt.Println("See you in Valhalla!")
```
