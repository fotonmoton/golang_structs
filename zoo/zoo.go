package zoo

// Private struct, visible inside this package only
type zookeeper struct {
	// Contains different field names with different types
	Name       string
	Surname    string
	Age        int
	Experience int
}

type Animal struct {
	Kind           string
	Name           string
	Number_of_legs int
	Predator       bool
}

type Cage struct {
	Size   string
	Color  string
	Number int
	Animal Animal
}

func (a *Animal) Food() {
	if a.Predator == true {
		println("it likes meat")
	} else {
		println("it likes apple")
	}

}

var Mihalich = zookeeper{"Mihail", "Ivanov", 56, 30}
