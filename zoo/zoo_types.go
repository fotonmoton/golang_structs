package zoo

// Animal struct
type Animal struct {
	Name   string
	Family string
}

// Cage struct
type Cage struct {
	Family  string
	Animals []Animal
}

// Zookeeper struct
type Zookeeper struct {
	Name string
}

// Zoo struct
type Zoo struct {
	Name      string
	Zookeeper Zookeeper
	Cages     []Cage
}
