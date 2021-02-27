package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type service interface {
	apiNavy()
}

type Satelite struct {
	name     string `json:"name"`
	distance string `json:"distance,omitempty"`
	message  string `json:"message,omitempty"`
}

type Position struct {
	x float32 `json:"x"`
	y float32 `json:"y"`
}

type SateliteResponse struct {
	position Position `json:"position"`
	messages string   `json:"message,omitempty"`
}

var Satelites []Satelite
var SatelitesResponse []SateliteResponse

func PostNaviHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Codificar")
		w.Header().Set("Content-Type", "application/json")
		reqBody, _ := ioutil.ReadAll(req.Body)
		var satelite Satelite
		json.Unmarshal(reqBody, &satelite)
		// update our global Articles array to include
		// our new Article
		var sateliteResponse SateliteResponse
		var position Position
		position.x = 100
		position.y = 75
		sateliteResponse.position = position
		sateliteResponse.messages = satelite.message
		SatelitesResponse := append(SatelitesResponse, sateliteResponse)
		fmt.Println(SatelitesResponse)
		json.NewEncoder(w).Encode(SatelitesResponse)
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func PostSplitNaviHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)
}

func handleRequests() {
	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/topsecret", PostNaviHandler).Methods("POST")
	r.HandleFunc("/topsecret_split/{satellite_name}", PostSplitNaviHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", r))
}

//func ApiNavy() {
func main() {
	handleRequests()
}
