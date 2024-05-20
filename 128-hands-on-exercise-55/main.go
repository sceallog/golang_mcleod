package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make   string
	model  string
	doors  int
	colour string
}

func main() {
	v1 := vehicle{
		engine: engine{electric: true},
		make:   "Tesla",
		model:  "Model 3",
		doors:  5,
		colour: "Red",
	}

	v2 := vehicle{
		engine: engine{electric: false},
		make:   "BMW",
		model:  "M3",
		doors:  5,
		colour: "Silver",
	}

	fmt.Printf("%#v\n", v1)
	fmt.Printf("%#v\n", v2)

	for _, v := range []vehicle{v1, v2} {
		fmt.Printf("%s %s has %d doors and is %s\n", v.make, v.model, v.doors, v.colour)
	}
}
