# simorgh
Simorgh (Phoenix in Persian) is a simple in-memory key/value database using radix tree.

![Image of Simorgh from Wikipedia](https://upload.wikimedia.org/wikipedia/commons/4/43/Phoenix-Fabelwesen.jpg)

Simorgh image from wikipedia

***

In-Memory Key/Value database based on radix tree with `Get` `Set` `Del` `Clr` commands.

### Download

```bash
git clone https://github.com/ahmdrz/simorgh
cd simorgh
make
```

And for install simrogh

```bash
sudo make install
```

### Running server

```bash
simrogh-server
```

Note that default port is 8080 and default protocol is tcp , you can pass `-port` and `-protocol` to `simorgh-server`

For more information :

```bash
simrogh-server --help
```

### Client

```bash
simorgh-client
```

For more information :

```bash
simrogh-client --help
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
