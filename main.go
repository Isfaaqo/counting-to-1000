// Declares that this is a group of functions and variables.
// Main shows where the start of the program is.
package main

import "fmt"

// Fmt is a library.

// Start of the program.
func main() {
	fmt.Println("Sup")

	// For assigning of a new variable.
	var fruit string = "Banana"
	fmt.Println(fruit)

	// Other way of assigning a new variable.
	name := "Jared"
	fmt.Println(name)

	// Reassign a new variable.
	name = "John"
	fmt.Println(name)

	// Assigning an integer, remember to add the bits needed.
	var age int32 = 18

	// Incrementing by one.
	age++

	// Always use float64 :)
	var balance float64 = 13.515

	// _ is used to void unused variables.
	_ = balance

	age = 17

	if age >= 18 {
		fmt.Println("Can drink")
	} else {
		fmt.Println("Cannot drink")
	}

	// Creating a list, defining the type and size.
	fruitBasket := make([]string, 0)
	// OR
	fruitBasket = []string{}

	// Appending something.
	fruitBasket = append(fruitBasket, "Oranges")

	// switch age {
	// case 0: fmt.Println("zero")
	// case 1: fmt.Println("one")
	// case 2: fmt.Println("two")
	// default:
	// }

	phonebook := map[string]int32{}

	phonebook["john"] = 32

	// ph, exists := phonebook["jill"]
	// if !exists {
	// handle
	// }

	seven()

}

func seven() int32 {
	defer fmt.Println("Some random thing")

	person := Person{
		Name: "John",
		Age:  82,
	}

	personList := []Person{}

	personList = append(personList, person)

	for _, person := range personList {
		person.SayName()
	}
	return 7
}

// Person represents a person object.
type Person struct {
	Name string // The name of the person
	Age  int32  // The age of the person
}

// SayName is calling function associated with Person.
func (p *Person) SayName() {
	fmt.Println(p.Name)
}
