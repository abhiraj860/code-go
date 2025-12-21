package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable] (s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo: ", SlicesIndex(s, "zoo"))
	_ = SlicesIndex[[]string, string](s, "zoo")
}