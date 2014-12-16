package main

import (
	"fmt"
)

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func main() {
	var fruit Fruit = Apple
	fmt.Println(fruit)
}
