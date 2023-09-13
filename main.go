package main

import (
	"fmt"

	"github.com/shxdxwraze/golang_structs/zoo"
)

func main() {
	// Create a zookeeper
	zk := zoo.Zookeeper{Name: "Yaroslav"}

	// Create some animals
	lion := zoo.Animal{Name: "Lion", Family: "Predator"}
	zebra := zoo.Animal{Name: "Zebra", Family: "Mammal"}
	gorilla := zoo.Animal{Name: "Gorilla", Family: "Mammal"}
	tiger := zoo.Animal{Name: "Tiger", Family: "Predator"}
	antelope := zoo.Animal{Name: "Antelope", Family: "Mammal"}

	animals := []zoo.Animal{lion, zebra, gorilla, tiger, antelope}

	// Create some cages
	predators := zoo.Cage{Family: "Predator"}
	mammals := zoo.Cage{Family: "Mammal"}

	cages := []zoo.Cage{predators, mammals}

	// Initiate zoo
	z := zoo.NewZoo("prjctr", zk)

	// act
	fmt.Print("Oh no all cages are empty\n")
	fmt.Printf("\nLet's call zookeeper %s to return animals:\n", zk.Name)

	// Let's collect animals
	for _, cage := range cages {
		fmt.Print("\n==============================================\n")
		fmt.Printf("Catching %s animals\n", cage.Family)
		fmt.Print("==============================================\n")
		for _, animal := range animals {
			if err := zk.CollectAnimals(&cage, animal); err != nil {
				fmt.Printf("Oh no %s why you %s\n", zk.Name, err)
			}
		}
		z.Cages = append(z.Cages, cage)
		fmt.Print("\n==============================================\n")
		fmt.Printf("Zookeeper %s caught all %s animals in their cage\n", zk.Name, cage.Family)
		fmt.Print("==============================================\n")
	}

	// Let's list animals in cages
	fmt.Print("\n==============================================\n")
	zk.ZkWorkingHard(z.Cages)
}
