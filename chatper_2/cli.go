package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "world", "a name to say hello")
var spanish bool


func init(){
	flag.BoolVar(&spanish, "spanish", false, "Use spanish")
	flag.BoolVar(&spanish, "s", false, "Use spanish")
}
func main() {
	flag.Parse()

	if spanish {
		fmt.Printf("Hola %s", *name)
	} else {
		fmt.Printf("Hello %s", *name)
	}
	fmt.Println()
}
