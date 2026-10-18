package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "apple banana apple orange banana apple"
	dict := make(map[string]int)
	for _, w := range strings.Fields(text) {
		dict[w]++
	}
	for k, v := range dict {
		fmt.Printf("Слово %q, Повторений: %d\n", k, v)
	}
}