package main

import (
	"fmt" //standard go library

	"github.com/jboursiquot/go-proverbs"
)

const location = "Remote" //type of constant is inferred and should not be necessarily be explicit

var name string //global variable which has only been declared

func main() {
	name = "Wilbrod" //variable has been assigned
	from := `Zimbabwe` //declaration and assignment done at the same time
	var n int = 2 

	var proverb = "Undefined"
	if p, err := proverbs.Nth(4); err == nil {
		proverb = p.Saying
	}

	fmt.Printf("Hello, fellow %s Gophers!\n", location)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock EST comes around, we'll know how to code in Go!\\n", n)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
}
