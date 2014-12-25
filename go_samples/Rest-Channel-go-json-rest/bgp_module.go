package main

import (
	"fmt"
	//"net"
	"time"
)

const ROUTE_FAMILY_IPv4 = "IPv4"
const ROUTE_FAMILY_IPv6 = "IPv6"

type Peer struct {
	//need to define a structure
	RemoteAs      uint32
	VersionNum    int
	RemoteAddress string
	protocol      *BgpProtocol
	Tables        map[string]*Table
}
type BgpProtocol struct {
	//need to define a structure
}
type Table struct {
	Destinations map[string]*Destination
}
type Destination struct {
	Paths []*Path
}
type Path struct {
	Nlri     string
	AsNum    int
	Nexthop  string
	Bestpath bool
}

func NewPeer(pref string) *Peer {
	peer := new(Peer)
	peer.RemoteAddress = pref
	peer.Tables = make(map[string]*Table)
	return peer
}
func NewTable() *Table {
	table := new(Table)
	table.Destinations = make(map[string]*Destination)
	return table
}
func NewDestination() *Destination {
	dest := new(Destination)
	dest.Paths = make([]*Path, 0)
	return dest
}
func NewPath(nlri string, asNum int, nexthop string, bestpath bool) *Path {
	path := &Path{
		Nlri:     nlri,
		AsNum:    asNum,
		Nexthop:  nexthop,
		Bestpath: bestpath,
	}
	return path
}

func (p *Peer) BGPRoute() {
	p1 := NewPath("10.10.10.0", 64000, "192.168.1.0", false)
	p2 := NewPath("10.10.10.0", 64001, "192.168.2.0", true)
	p3 := NewPath("20.20.20.0", 64000, "192.168.3.0", true)
	p4 := NewPath("20.20.20.0", 64001, "192.168.4.0", false)
	p5 := NewPath("fe80:1234:1234:5667:8967:af12:1203:33a1", 64000, "192.168.5.0", false)
	p6 := NewPath("fe80:1234:1234:5667:8967:af12:1203:33a1", 64001, "192.168.6.0", false)
	p7 := NewPath("fe80:1234:1234:5667:8967:af12:8912:1023", 64000, "192.168.7.0", true)
	p8 := NewPath("fe80:1234:1234:5667:8967:af12:8912:1023", 64001, "192.168.8.0", true)
	dest1 := NewDestination()
	dest2 := NewDestination()
	dest3 := NewDestination()
	dest4 := NewDestination()
	dest1.Paths = []*Path{p1, p2}
	dest2.Paths = []*Path{p3, p4}
	dest3.Paths = []*Path{p5, p6}
	dest4.Paths = []*Path{p7, p8}
	tab4 := NewTable()
	tab6 := NewTable()
	tab4.Destinations["10.10.10.0"] = dest1
	tab4.Destinations["20.20.20.0"] = dest2
	tab6.Destinations["fe80:1234:1234:5667:8967:af12:1203:33a1"] = dest3
	tab6.Destinations["fe80:1234:1234:5667:8967:af12:8912:1023"] = dest4

	p.Tables[ROUTE_FAMILY_IPv4] = tab4
	p.Tables[ROUTE_FAMILY_IPv6] = tab6

	fmt.Println(p.Tables)
}
func (p *Peer) BGPRouteMod() {
	delete(p.Tables, ROUTE_FAMILY_IPv6)
}

func (p *Peer) GetRoute(trig *Trigger) {
	for {
		select {
		case pref := <-trig.chPref:
			fmt.Println("exec get route!!! [ ", pref, " ]")
			trig.chRoute <- p
		default:
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		}
	}
}
