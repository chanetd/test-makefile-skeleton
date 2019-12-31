package main

import (
	"fmt"

	"github.com/namsral/flag"
)

var flagX = flag.String("x", "something", "something")

func main() {
	flag.Parse()
	fmt.Println(*flagX)
}
