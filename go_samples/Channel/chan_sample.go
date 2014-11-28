// chan_sample.go
package main

import (
	"fmt"
	"time"
)

const (
	_ = iota
	COLOR_RED
	COLOR_BURE
	COLOR_BLACK
)

type Car struct {
	Name      string
	Color     map[int]string
	SpeedBase int
	Gus       int
}

func NewCar() *Car {
	car := &Car{}
	car.Color = make(map[int]string)
	return car
}
func main() {
	car := NewCar()
	car.Name = "stinglay"
	car.Color[COLOR_BLACK] = "black"
	car.SpeedBase = 30
	car.Gus = 100

	fmt.Println("# Start #")
	go Calculate(car)
	time.Sleep(10 * time.Second)
	fmt.Println("# End #")
}
func Calculate(car *Car) {
	for i := 0; i < 1000; i++ {
		car.SpeedBase += 1
		car.Gus -= 2
		time.Sleep(1 * time.Second)
		fmt.Println(car)
	}
}
