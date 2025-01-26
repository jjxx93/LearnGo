package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{7, 9, 8}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
