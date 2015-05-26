# xtime

Go package `xtime` returns atime, ctime and mtime

## Desciption

> [How can I get file's ctime,mtime,atime use golang and change them?](http://stackoverflow.com/questions/20875336/how-can-i-get-a-files-ctime-atime-mtime-and-change-them-using-golang) - Stack Overflow

Use [b4b4r07/xtime](https://github.com/b4b4r07/xtime), it would solve.

## Usage


```go
package main

import (
	"github.com/b4b4r07/xtime"
	"log"
	"os"
)

func main() {
	fi, err := os.Stat("file")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n%s\n%s\n%s\n",
		xtime.Atime(fi),
		xtime.Ctime(fi),
		xtime.Mtime(fi),
	)
}
```

## Installation

```sh
go get github.com/b4b4r07/xtime
```

## Thanks

It was heavily inspired by [djherbis/atime](https://github.com/djherbis/atime).

## License

[MIT](https://raw.githubusercontent.com/b4b4r07/dotfiles/master/doc/LICENSE-MIT.txt)

## Author

[BABAROT](http://tellme.tokyo) a.k.a. b4b4r07
