package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	S := sc.Text()
	if strings.HasSuffix(S, "er") {
		fmt.Println("er")
	} else if strings.HasSuffix(S, "ist") {
		fmt.Println("ist")
	}
}
