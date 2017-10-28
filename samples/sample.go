package main

import "fmt"

func main() {

	m := make([]int, 3, 5)
	n := make([]int, 4)
	m[0] = 1
	m = append(m, 4)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))




	map1 := make(map[string]int)
	map1["tetest"] = 1

	fmt.Println(map1)

}
