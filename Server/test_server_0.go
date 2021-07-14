package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Numbers struct {
	First uint16 `json:"first"`
	Second uint16 `json:"second"`
}

var Response struct {
	Numbers []uint16 `json:"numbers"`
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/fibo", showList).Methods("GET")
	http.ListenAndServe(":80", r)
}

func showList(w http.ResponseWriter, r *http.Request){
	var newNumbers Numbers
	data, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal(data, &newNumbers)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}
	w.Header().Set("Content-Type", "application/json")
	fibonacci(&newNumbers)
	json.NewEncoder(w).Encode(Response.Numbers)
}

func fibonacci(newNumbers *Numbers){
	allNumbers:=make([]uint16, newNumbers.Second)
	allNumbers[0] = 0
	allNumbers[1] = 1
	for i:=2;i<len(allNumbers);i++{
		allNumbers[i] = allNumbers[i-1] + allNumbers[i-2]
	}
	Response.Numbers = allNumbers[newNumbers.First:]
}
