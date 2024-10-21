package main

import (
	"fmt"
)

type Person struct {
	FistName string
	LastName string
	Age      int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, %d", p.FistName, p.LastName, p.Age)
}
func main() {
	p := Person{"John", "Doe", 25}
	fmt.Println(p)
}
