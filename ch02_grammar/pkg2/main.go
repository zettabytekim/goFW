package main

import (
	"GoWebFW/ch02_grammar/lib"
	"fmt"
)

var v rune

func init() {
	v = '1'
}

func main() {
	fmt.Println(lib.IsDigit(v))
}

// IsDigit
func IsDigit(c int32) bool {
	return '0' < c && c <= '9'
}
