# Simorgh
> It was my very very old golang-codes. So it maybe has a lot of bugs or bad written codes!

### Under Construction :)

Simorgh (Phoenix in Persian) is a simple in-memory key/value database using radix tree. Protected by SRP authentication algorithm.

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

*Default Username is :* `simorgh`
*Default Password is :* `simorgh`

Client program :

```
Username > simorgh
Password > 
< set a=b
> OK OK
< get a
> OK b
< get b=c
> ERROR ! UNDEFINED
< get b
> ERROR ! UNDEFINED
< set b=c
> OK OK
< get b
> OK c
< del b
> OK OK
< clr
> OK MEMORY CLEARED (1)
< \q
bye
```

### API

You can use `simorgh` in your Golang program.
After installation you can do some code like :

```golang
package main

import (
    "fmt"
    "github.com/ahmdrz/simorgh/driver"    
)

func main() {
    si, err := simorgh.New("localhost:8080","simorgh","simorgh","tcp")
	if err != nil {
		panic(err)
	}
	defer si.Close()

    fmt.Println(si.Set("a","b"))
    fmt.Println(si.Get("a"))
    fmt.Println(si.Clr())
}
```

### TODO

- [x] Password authentication.
- [ ] Save configuration file in encrypted text file.
- [ ] Improve Simorgh Cli.
- [ ] Build Simorgh Golang library.
- [ ] Test with heavy dataset.
- [ ] Improve Simorgh base architecture.
- [ ] Make some test files and pass it to Travis CI.

### Contribute

`Radix Tree` is forked from [armon](https://github.com/armon/go-radix). 

I'm not good in data structures , So I will happy if anyone give me suggestions and improve my code.

***

Build with :heart: in Iran.
