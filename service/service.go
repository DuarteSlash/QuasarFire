package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type service interface {
	apiNavy()
}

type Satelite struct {
	Name     string  `json:"name,string"`
	Distance float32 `json:"distance,string"`
	Message  string  `json:"message,,string"`
}

type Position struct {
	X float32 `json:"x,string"`
	Y float32 `json:"y,string"`
}

type SateliteResponse struct {
	Position Position `json:"position,string"`
	Messages string   `json:"message,string"`
}

var Satelites []Satelite
var SatelitesResponse []SateliteResponse = []SateliteResponse{}

func PostNaviHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Codificar")

		//reqBody, _ := ioutil.ReadAll(req.Body)
		var satelite Satelite
		json.NewDecoder(req.Body).Decode(&Satelites)
		fmt.Println(satelite)

		//json.Unmarshal(reqBody, &satelite)
		// update our global Articles array to include
		// our new Article
		var sateliteResponse SateliteResponse
		var position Position

		position.X = 100
		position.Y = 75
		sateliteResponse.Position = position
		sateliteResponse.Messages = "este es un mensaje secreto"
		SatelitesResponse := append(SatelitesResponse, sateliteResponse)
		fmt.Println(position.X)
		fmt.Println(position.Y)
		fmt.Println(SatelitesResponse)
		w.Header().Set("Content-Type", "application/json")
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
