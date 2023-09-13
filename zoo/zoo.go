package zoo

import (
	"fmt"
	"reflect"
)

// NewZoo returns zoo
func NewZoo(name string, zk Zookeeper) *Zoo {
	return &Zoo{
		Name:      name,
		Zookeeper: zk,
	}
}

// CheckCageCompatibility return error if zookeeper add animals with different behaviour in same cage
func (c *Cage) CheckCageCompatibility(animal Animal) error {
	if !reflect.DeepEqual(c.Family, animal.Family) {
		return fmt.Errorf("trying to add an %s to a cage for %s", animal.Name, c.Family)
	}
	return nil
}

// CollectAnimals zookeeper collects animals to the cage
func (z *Zookeeper) CollectAnimals(cage *Cage, animal Animal) error {
	if err := cage.CheckCageCompatibility(animal); err != nil {
		return err
	}

	cage.Animals = append(cage.Animals, animal)
	fmt.Printf("Zookeper %s added %s to %s cage\n", z.Name, animal.Name, cage.Family)

	return nil
}

// ZkWorkingHard shows collected animals by Zookeeper
func (z *Zookeeper) ZkWorkingHard(cages []Cage) error {
	for _, cage := range cages {
		for _, animal := range cage.Animals {
			fmt.Printf("In cage %s we have %s is from %s family\n", cage.Family, animal.Name, animal.Family)
		}
	}

	return nil
}
