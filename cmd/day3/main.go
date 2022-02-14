package main

import (
	"fmt"

	_ "github.com/pallat/hello/cmd/day3/effect"
)

func init() {
	fmt.Println("before everything")
}
func init() {
	fmt.Println("before everything 1")
}

func main() {
	fmt.Println("Hello, world")
}
