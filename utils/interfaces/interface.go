package main

import "fmt"

type animal interface {
	move() string
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

type dog struct{}
type cat struct{}
type fish struct{}

func (dog) move() string {
	return "Im a dog"
}

func (cat) move() string {
	return "Im a cat"
}

func (fish) move() string {
	return "Im a fish"
}

func main() {
	d := dog{}
	c := cat{}
	f := fish{}
	moveAnimal(d)
	moveAnimal(c)
	moveAnimal(f)
}
