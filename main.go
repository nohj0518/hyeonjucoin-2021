package main

import (
	"fmt"

	"github.com/nohj0518/nomadcoin/person"
)

func main() {
	fmt.Println("3.7 Structs with Pointers!")
	hyeonju := person.Person{} // Create new person Hyeonju
	hyeonju.SetDetails("hyeonju", 18)
	fmt.Println("Main 현주", hyeonju)
}
