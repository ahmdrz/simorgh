# simorgh
Simorgh (Phoenix in Persian) is a simple in-memory key/value database using radix tree.

***

In-Memory Key/Value database based on radix tree with `Get` `Set` `Del` `Clr` commands.

### Download

```bash
go get github.com/ahmdrz/simorgh
```

### Client

example of client connection in `client` directory

```bash
cd $GOPATH/src/github.com/ahmdrz/simorgh/client
go build -i
./client -port=8080 -protocol=tcp -address=localhost
```

### Server

Simorgh server is in `server` directory. 

```bash
cd $GOPATH/src/github.com/ahmdrz/simorgh/server
go build -i
./server -port=8080 -protocol=tcp
```

```bash

```

### TODO

- [ ] Password authentication.
- [ ] Improve Simorgh Cli.
- [ ] Build Simorgh Golang library.
- [ ] Test with heavy dataset.
- [ ] Improve Simorgh base architecture.
- [ ] Make some test files and pass it to Travis CI.

### Contribute

I'm not good in data structures , So I will happy if anyone give me suggestions and improve my code.

***

Build with :heart: in Iran.
