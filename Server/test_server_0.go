package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Numbers struct {
	First int32 `json:"first"`
	Second int32 `json:"second"`
}

var Response struct {
	Numbers []int32 `json:"numbers"`
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/fibo", showList).Methods("GET")
	http.ListenAndServe(":80", r)
}

func showList(w http.ResponseWriter, r *http.Request){
	var newNumbers Numbers
	json.NewDecoder(r.Body).Decode(&newNumbers)
	w.Header().Set("Content-Type", "application/json")
	fibonacci(&newNumbers)
	json.NewEncoder(w).Encode(Response.Numbers)
}

func fibonacci(newNumbers *Numbers){
	allNumbers:=make([]int32, newNumbers.Second)
	allNumbers[0] = 0
	allNumbers[1] = 1
	for i:=2;i<len(allNumbers);i++{
		allNumbers[i] = allNumbers[i-1] + allNumbers[i-2]
	}
	Response.Numbers = allNumbers[newNumbers.First:]
}
