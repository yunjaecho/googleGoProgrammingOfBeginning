package chapter31

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func init() {
	fmt.Println("=== JSON Publisher ===")
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:8181", nil)
}

func serveRest(w http.ResponseWriter, r *http.Request){
	response, err := publishJson()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func publishJson()([] byte, error) {
	cars := make(map[string]int)
	cars["2Door"] = 20
	cars["4Door"] = 32

	trucks := make(map[string]int)
	trucks["1Door"] = 14
	trucks["2Door"] = 4

	d := Data{cars, trucks}
	p := Payload{d}

	//return json.MarshalIndent(p,  "Root", " ")
	return json.MarshalIndent(p,  "", " ")
}

type Payload struct {
	Envelope Data
}

type Data struct {
	CarsByDoor Cars
	TrucksByTon Trucks
}

type Cars map[string] int
type Trucks map[string] int