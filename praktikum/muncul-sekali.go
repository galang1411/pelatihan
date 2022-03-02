package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {

	check := map[string]int{}
	res := make([]int, 0)
	for _, val := range angka {
		check[string(val)] += 1
	}
	for letter, div := range check {
		frina, _ := strconv.Atoi(letter)
		if div == 1 {
			res = append(res, frina)
		}
	}
	return res
}

func main() {
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("1122"))
}
