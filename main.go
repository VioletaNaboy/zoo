package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Animal struct {
	Name string
}

type Cage struct {
	Animal *Animal
}

type Zookeeper struct {
	Name  string
	Cages []*Cage
}

func (z *Zookeeper) CollectAnimal(a *Animal)  {
	for i, cage := range z.Cages {
		if cage.Animal == nil {
			cage.Animal = a
			fmt.Printf("%s collected %s and put it in cage %d\n", z.Name, a.Name, i+1)
			return
		}
	}
	fmt.Printf("No empty cage found for %s\n", a.Name)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	animal1 := Animal{Name: "Lion"}
	animal2 := Animal{Name: "Tiger"}
	animal3 := Animal{Name: "Elephant"}
	animal4 := Animal{Name: "Giraffe"}
	animal5 := Animal{Name: "Zebra"}

	cage1 := &Cage{}
	cage2 := &Cage{}
	cage3 := &Cage{}
	cage4 := &Cage{}
	cage5 := &Cage{}

	zookeeper := Zookeeper{
		Name: "John",
	    Cages: []*Cage{cage1, cage2, cage3, cage4, cage5},
	}

	animals := []*Animal{&animal1, &animal2, &animal3, &animal4, &animal5}
	escapedCount := rand.Intn(len(animals)) + 1
	escapedAnimals := animals[:escapedCount]

	fmt.Println("Escaped animals:")
	for _, a := range escapedAnimals {
		fmt.Println(a.Name)
	}

	for _, a := range escapedAnimals {
		zookeeper.CollectAnimal(a)
	}

	fmt.Println("\nFinal state of cages:")
	for i, cage := range zookeeper.Cages  {
		if cage.Animal != nil {
			fmt.Printf("Cage %d: %s\n", i+1, cage.Animal.Name)
		} else {
			fmt.Printf("Cage %d: Empty\n", i+1)
		}
	}
}
