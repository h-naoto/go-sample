// rest_sample.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const BASE_VERSION = "/v1"
const RES_RIB = "/local_rib"
const RES_RIB_BEST = "/local_rib_best"
const REST_PORT = "3000"
const PARAM_PREFIX = "prefix"

type Trigger struct {
	chPref  chan string
	chRoute chan *Peer
}

func NewTrigger() *Trigger {
	trig := new(Trigger)
	trig.chPref = make(chan string, 1)
	trig.chRoute = make(chan *Peer, 1)
	return trig
}

func (t *Trigger) StartRestService() {
	neighbor_rib := BASE_VERSION + RES_RIB
	neighbor_rib_best := BASE_VERSION + RES_RIB_BEST
	r := mux.NewRouter()
	// set URLs
	//r.HandleFunc(neighbor_rib+"/{"+PARAM_PREFIX+"}/state", t.GetLocalRib).Methods("GET")
	//r.HandleFunc(neighbor_rib_best+"/{"+PARAM_PREFIX+"}/state", t.GetLocalRibBest).Methods("GET")
	tmp := neighbor_rib + "/{" + PARAM_PREFIX + "}/state"
	tmp1 := neighbor_rib_best + "/{" + PARAM_PREFIX + "}/state"
	r.HandleFunc(tmp, t.GetLocalRib).Methods("GET")
	r.HandleFunc(tmp1, t.GetLocalRibBest).Methods("GET")

	// Handler when not found url
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	http.Handle("/", r)

	http.ListenAndServe(":"+REST_PORT, nil)
}

func (t *Trigger) GetLocalRib(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	remoteAddr, _ := params[PARAM_PREFIX]
	fmt.Println(remoteAddr)
	t.chPref <- remoteAddr
	peer := <-t.chRoute
	fmt.Println("get peer!!! [ ", peer, " ]")
	res, _ := json.MarshalIndent(peer, "", "  ")
	w.Write(res)
}

func (t *Trigger) GetLocalRibBest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	remoteAddr, _ := params[PARAM_PREFIX]
	fmt.Println(remoteAddr)
	t.chPref <- remoteAddr
	peer := <-t.chRoute
	fmt.Println("get peer!!! [ ", peer, " ]")
	res, _ := json.MarshalIndent(peer, "", "  ")
	w.Write(res)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
