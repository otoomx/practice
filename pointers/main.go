package main

import "log"

type test interface {

}

type foo struct {
	Name string
}

func main() {

	myName := foo{Name:"mike"}
	log.Printf("Name: %s\n", myName.Name)
	nopointerDemo(myName)
	log.Printf("Name: %s\n", myName.Name)
	pointerDemo(&myName)
	log.Printf("Name: %s\n", myName.Name)


}


func nopointerDemo(name foo){
	name.Name = "bar"
	log.Printf("Name: %s\n", name.Name)
}

func pointerDemo(name *foo){
	name.Name = "bar"
	log.Printf("Name: %s\n", name.Name)

}


