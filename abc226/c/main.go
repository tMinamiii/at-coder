package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var n int64
	sc.Scan()
	nstr := sc.Text()
	n, _ = strconv.ParseInt(nstr, 10, 64)

	input := make(map[string]int)
	for i := int64(0); i < n; i++ {
		sc.Scan()
		line := sc.Text()
		input[line]++

	}
	fmt.Println(len(input))
}
