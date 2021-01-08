package main

import (
	"fmt"
	"messages"
)

func main() {
	var amount messages.Value = 55
	fmt.Printf("Hello Marcos %d", amount.SumTen())
}
