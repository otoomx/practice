package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {

	p1 := Person{"Mike", "OToole", 1, 1}
	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Println(string(bs))
}
