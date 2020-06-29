package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	var s string
	if len(args) > 1 {
		s = strings.Join(args[1:], " ")
	}
	fmt.Println(s)
}
