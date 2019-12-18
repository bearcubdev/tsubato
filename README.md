# Tsubato つばと

[![GoDoc](https://godoc.org/github.com/bearcubdev/tsubato?status.svg)](https://godoc.org/github.com/bearcubdev/tsubato)

Tsubato is a 4chan API query library. It follows the japanese teahouse philosophy in that there's not much here.

## Install

```
go get -u github.com/bearcubdev/tsubato
```

## Features:

* **Satisfies 4chan API usage policy**
  * 1 request / second speed limit
  * Caches last request time and sends the if-modified-since header for you
* ... plus standard features
  * Request a list of boards
  * Request a catalog by board
  * Request a thread by ID

## Example usage:

```go
import (
	"fmt"
	"time"

	"github.com/bearcubdev/tsubato"
)

func main() {
	board, err := tsubato.GetBoard("g")
	fmt.Println(board.Title, err) // Prints "Technology <nil>"

	_, err = tsubato.GetCatalog("g")
	fmt.Println(err) // Prints "https://a.4cdn.org/g/catalog.json requested too recently"

	// So we wait 2s between requests to ensure we don't hit the timer.
	time.Sleep(time.Second * 2)

	catalog, err := tsubato.GetCatalog("g")
	fmt.Println(len(catalog.Pages), err) // Prints "11 nil"
}
```
