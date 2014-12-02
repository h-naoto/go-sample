// rest_sample.go
package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
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
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	rib_url := BASE_VERSION + RES_RIB
	reb_best_url := BASE_VERSION + RES_RIB_BEST
	// set URLs
	handler.SetRoutes(
		&rest.Route{"GET", rib_url + "/#" + PARAM_PREFIX, t.GetLocalRib},
		&rest.Route{"GET", reb_best_url + "/#" + PARAM_PREFIX, t.GetLocalRibBest},
	)
	http.ListenAndServe(":"+REST_PORT, &handler)
}

func (t *Trigger) GetLocalRib(w rest.ResponseWriter, r *rest.Request) {
	peerPrefix := r.PathParam(PARAM_PREFIX)
	fmt.Println(peerPrefix)
	t.chPref <- peerPrefix
	peer := <-t.chRoute
	fmt.Println("get peer!!! [ ", peer, " ]")
	w.WriteJson(peer)
}

func (t *Trigger) GetLocalRibBest(w rest.ResponseWriter, r *rest.Request) {
	peerPrefix := r.PathParam(PARAM_PREFIX)
	fmt.Println(peerPrefix)
	t.chPref <- peerPrefix
	peer := <-t.chRoute
	fmt.Println("get peer!!! [ ", peer, " ]")
	w.WriteJson(peer)
}
