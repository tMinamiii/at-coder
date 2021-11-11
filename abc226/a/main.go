package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Scan(&input)
	fmt.Printf("%d\n", int(math.Round(input)))
}
