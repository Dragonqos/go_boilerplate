# GO-Boilerplate

## Using with Docker
Start production enviroment
```console
$ not ready for production
```

Start development enviroment 
```console
$ docker network create shared_go_bp
$ docker-compose up
```

Attach to bash
```console
$ docker-compose exec go_go sh
$ de go_go
```

## Run tests
Enter go_go container
```console
$ de go_go
```
Local directory mode is when go test is called inside a directory, without any package arguments supplied. Running `go test` in `.../yourProject/model` will compile and run all the tests in that directory 
Package list mode is when go test is called with package/path arguments, such as `go test users` or `go test ./testFiles/` .
```console
$ cd /api/core/cipher
$ go test
// or run all tests from root
$ cd /api
$ go test ./...
```

development: http://go-local.com:8200/api/
production:  not ready 


## Example
After running docker-compose open:

development: http://go-local.com:8200/api/
production:  not ready 

## Without docker
```bash
$ make serve
``` 
You may need to execute `go mod download` in `src` folder first
