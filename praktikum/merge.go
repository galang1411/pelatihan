package main

import (
	"fmt"
)

func ArrayMerge(arrayA []string, arrayB []string) []string {

	check := map[string]bool{}
	d := append(arrayA, arrayB...)
	res := make([]string, 0)
	for _, val := range d {
		check[val] = true
	}
	for letter, _ := range check {
		res = append(res, letter)
	}
	return res
}

func main() {

	fmt.Println(ArrayMerge([]string{"kazuya", "jin", "lee"}, []string{"kazuya", "feng"}))
	fmt.Println(ArrayMerge([]string{"lee", "jin"}, []string{"kazuya", "panda"}))
}
