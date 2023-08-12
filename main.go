package main

import (
	"fmt"
	"sort"
)

func main() {

	frutas := []string{"3uva", "7pessego", "6banana", "7morango", "8melancia", "4kiwi", "4maça"}

	organizarFrut := func(i, j int) bool {
		return len(frutas[i]) < len(frutas[j])
	}

	sort.Slice(frutas, organizarFrut)

	fmt.Println(frutas)
}