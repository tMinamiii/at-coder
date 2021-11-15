package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()

	chars := make(map[rune]struct{})
	for _, char := range line {
		chars[char] = struct{}{}
	}
	switch len(chars) {
	case 3:
		fmt.Println(6)
	case 2:
		fmt.Println(3)
	case 1:
		fmt.Println(1)
	}
}
