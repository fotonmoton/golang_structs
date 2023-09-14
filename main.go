package main

import (
	"fmt"
	"structs/zoo"
)

/*
0. Declaration
1. Different types
3. Creating instances, instantiation
4. Zero values
5. Consecutive fields with same type
6. User defined type. Structs are the only way to create concrete user-defined types in Golang.
7. Struct Visibility
8. Fields Visibility
9. Embedding
10. Recursive (almost)
11. Anonymous/unnamed structs
12. Structural comparison
13. methods or pointer and value receivers
14. Tagging
15. Implicit dereferencing
16. Assignment
*/

func main() {

	cage1 := zoo.Cage{Size: "10x10", Color: "green", Number: 1, Animal: zoo.Animal{Kind: "Leon", Name: "Pusic", Number_of_legs: 4, Predator: true}}
	cages := []zoo.Cage{cage1, {Size: "20x20", Color: "red", Number: 2, Animal: zoo.Animal{Kind: "mouse", Name: "Miki", Number_of_legs: 4, Predator: false}}}

	leon := zoo.Animal{Kind: "Elephant", Name: "White", Number_of_legs: 4, Predator: false}
	cages = append(cages, zoo.Cage{Size: "10x10", Color: "green", Number: 1, Animal: leon})

	fmt.Printf("On duty today: %s, with %d years of experience\n", zoo.Mihalich.Name, zoo.Mihalich.Experience)
	fmt.Printf("We have a %v cages\n", len(cages))
	for i := 0; i < len(cages); i++ {
		fmt.Printf("In cage: %v cages, we have %s, named %s. And ", i+1, cages[i].Animal.Kind, cages[i].Animal.Name)
		cages[i].Animal.Food()
	}

}
