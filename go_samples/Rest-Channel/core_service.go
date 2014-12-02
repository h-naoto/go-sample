// core_service
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	CoreService()
}

func CoreService() {
	fmt.Println("Hello core service")
	pr := make(map[string]*Peer)
	pr["172.16.10.0"] = NewPeer("172.16.10.0")
	pr["172.16.10.0"].BGPRoute()

	trig := NewTrigger()
	go pr["172.16.10.0"].GetRoute(trig)
	go trig.StartRestService()
	time.Sleep(10 * time.Second)
	fmt.Println("modified route!!")
	pr["172.16.10.0"].BGPRouteMod()
	time.Sleep(1000 * time.Second)
}
