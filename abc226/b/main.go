package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var n int
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	input := make(map[string]int)
	for i := 0; i < n; i++ {
		sc.Scan()
		line := sc.Text()
		input[line]++

	}
	fmt.Println(len(input))
}
