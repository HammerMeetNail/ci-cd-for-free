package main

import (
	intutils "app/pkg"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Arg must be a number that can convert to float64")
	}
	square := intutils.Square(num)
	fmt.Printf("%v\n", square)
}
