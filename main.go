package main

import (
	"fmt"
	"structs/zoo"
)

func main() {

	myZoo := zoo.MakeZoo(108)

	fmt.Printf("We have a great zoo, it has %d cages!\n", len(myZoo.Cages))
	for _, cage := range myZoo.Cages {
		fmt.Printf("In %d cage we have a %d animals.\nNow I will have a little talk about them.\n", cage.Number, len(cage.Animals))
		for _, animal := range cage.Animals {
			fmt.Printf("%s:\n", animal.Name)
			fmt.Printf("- Population size: %s\n", animal.Characteristics.EstimatedPopulationSize)
			fmt.Printf("- Weight: %s\n", animal.Characteristics.Weight)
			fmt.Printf("- Height: %s\n", animal.Characteristics.Height)
			fmt.Printf("- Top speed: %s\n", animal.Characteristics.TopSpeed)
			fmt.Printf("-----------------------------------\n")
		}
	}
}
