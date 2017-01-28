# simorgh
Simorgh (Phoenix in Persian) is a simple in-memory key/value database using radix tree.

![Image of Simorgh from Wikipedia](https://upload.wikimedia.org/wikipedia/commons/4/43/Phoenix-Fabelwesen.jpg)

Simorgh image from wikipedia

***

In-Memory Key/Value database based on radix tree with `Get` `Set` `Del` `Clr` commands.

### Download

```bash
go get github.com/ahmdrz/simorgh
```

And for start simorgh server

```bash
cd $GOPATH/bin
./simorgh -port=8080 -protocol=tcp
```

### Client

example of client connection in `client` directory

```bash
cd $GOPATH/src/github.com/ahmdrz/simorgh/client
go build -i
./client -port=8080 -protocol=tcp -address=localhost
```

Client program :

```
> set ahmad=reza
< ahmad = reza
> get ahmad
< ahmad = reza
> get test
< test = UNDEFINED
> del ahmad
< ahmad REMOVED
.
.
.
```

### TODO

- [ ] Password authentication.
- [ ] Improve Simorgh Cli.
- [ ] Build Simorgh Golang library.
- [ ] Test with heavy dataset.
- [ ] Improve Simorgh base architecture.
- [ ] Make some test files and pass it to Travis CI.

### Contribute

`Radix Tree` is forked from [arman](https://github.com/armon/go-radix). 

I'm not good in data structures , So I will happy if anyone give me suggestions and improve my code.

***

Build with :heart: in Iran.
