// chan_sample.go
package main

import (
	"fmt"
	//	"launchpad.net/tomb"
	"time"
)

type Car struct {
	Name    string
	Color   string
	Speed   int
	UpSpeed int
	Gas     int
	GasCons int
}

func NewCar(name string, color string, bspeed int, cas int) *Car {
	car := &Car{
		Name:    name,
		Color:   color,
		Speed:   bspeed,
		UpSpeed: 1,
		Gas:     cas,
		GasCons: 2,
	}
	return car
}

//var Tomb tomb.Tomb

func main() {
	car := NewCar("stinglay", "black", 30, 30)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	fmt.Println("# Start #")
	go SpeedUp(car, ch1)
	go OutgoGas(car, ch2)
	for {
		select {
		case speed := <-ch1:
			fmt.Println("Car Speed -> ", speed, "km")
		case gas := <-ch2:
			if gas < 0 {
				return
			}
			fmt.Println("Car Gas -> ", gas, "L")
		default:
			fmt.Println("---------------------------------")
			time.Sleep(1 * time.Second)
		}
	}
	//time.Sleep(10 * time.Second)
	fmt.Println("# End #")
}
func SpeedUp(car *Car, ch1 chan int) {
	for {
		car.Speed += car.UpSpeed
		ch1 <- car.Speed
		time.Sleep(1 * time.Second)
		//fmt.Println("Car Speed -> ", car.Speed, "km")
	}
}
func OutgoGas(car *Car, ch2 chan int) {
	for {
		car.Gas -= car.GasCons
		ch2 <- car.Gas
		time.Sleep(1 * time.Second)
		//fmt.Println("Car Speed -> ", car.Gas, "L")
	}
}
