package main

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
