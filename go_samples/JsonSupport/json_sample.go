package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := "{RemoteAs:0,VersionNum:0,RemoteAddress:172.16.10.0,Tables:{IPv4:{Destinations:{10.10.10.0:{Paths:[{Nlri:10.10.10.0,AsNum:64000,Nexthop:192.168.1.0,Bestpath:false},{Nlri:10.10.10.0,AsNum:64001,Nexthop:192.168.2.0,Bestpath:true}]},20.20.20.0:{Paths:[{Nlri:20.20.20.0,AsNum:64000,Nexthop:192.168.3.0,Bestpath:true},{Nlri:20.20.20.0,AsNum:64001,Nexthop:192.168.4.0,Bestpath:false}]}}}}}"
	json.Encoder
	fmt.Println(jsonStr)
}
