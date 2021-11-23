package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	n := sc.Text()
	N, _ := strconv.Atoi(n)
	sc.Scan()
	rawInput := sc.Text()
	inputStr := strings.Split(rawInput, " ")
	count := 0
	for _, in := range inputStr {
		v, _ := strconv.Atoi(in)
		if hasArea(v) {
			count++
		}
	}
	fmt.Println(N - count)
}

func hasArea(v int) bool {
	for a := 1; a <= v; a++ {
		for b := 1; b <= v; b++ {
			if area(a, b) == v {
				return true
			}
		}
	}
	return false
}

func area(a, b int) int {
	return 4*a*b + 3*a + 3*b
}
