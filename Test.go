package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 100
	fmt.Printf("%v, %T", strconv.Itoa(i), strconv.Itoa(i))
}
