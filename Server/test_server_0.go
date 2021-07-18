package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Numbers struct {
	First uint16 `json:"first"`
	Second uint16 `json:"second"`
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/fibo", showList)
	err:= http.ListenAndServe(":80", r)
	if err != nil{
		fmt.Printf("Error listening: %v", err)
	}
}

func showList(w http.ResponseWriter, r *http.Request){
	var newNumbers Numbers
	data, readingErr:= ioutil.ReadAll(r.Body)
	if readingErr != nil{
		w.WriteHeader(400)
		w.Write([]byte("Empty response"))
		return
	}
	jsonErr := json.Unmarshal(data, &newNumbers)
	if jsonErr != nil{
		w.WriteHeader(400)
		w.Write([]byte("The numbers must be integers and positive"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(fibonacci(&newNumbers))
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("Something went wrong"))
		return
	}
}

func fibonacci(newNumbers *Numbers) []uint16 {
	allNumbers:=make([]uint16, newNumbers.Second)
	allNumbers[0] = 0
	allNumbers[1] = 1
	for i:=2;i<len(allNumbers);i++{
		allNumbers[i] = allNumbers[i-1] + allNumbers[i-2]
	}
	return(allNumbers[newNumbers.First:])
}
