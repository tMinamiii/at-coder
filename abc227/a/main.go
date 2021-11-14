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
	line := sc.Text()
	s := strings.Split(line, " ")
	N, _ := strconv.Atoi(s[0])
	K, _ := strconv.Atoi(s[1])
	A, _ := strconv.Atoi(s[2])
	n := A
	for i := 1; i < K; i++ {
		n++
		if n > N {
			n = 1
		}
	}
	fmt.Println(n)
}
