package main

import "fmt"

type Person struct {
	First string
	Middle string
	Last string
}
type DoubleZero struct {
	Person
	LicenseToKill bool
}
func main() {
	p1 := Person{"mike", "james", "otoole"}
	p2 := DoubleZero{Person{First:"fii",Middle:"dfd",Last:"sdf",},false}

	fmt.Println(p1)
	fmt.Println(p2.First)
	fmt.Println(p1.First)
}
