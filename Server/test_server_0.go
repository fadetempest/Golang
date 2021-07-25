package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Numbers struct {
	First uint16 `json:"first"`
	Second uint16 `json:"second"`
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/fibo", showList)
	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	go func(){
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Printf("Error listening: %v", err)
		}
	}()
	ch := make(chan os.Signal,1)
	ctx:= context.Background()
	signal.Notify(ch,syscall.SIGINT, syscall.SIGTERM)
	<-ch // заблочимся на этом моменте до комбинации клавиш
	srvErr:=srv.Shutdown(ctx) // у сервера метод какой-то такой есть посмотри его и соответсвенно контекст создай обычный заранее
	if srvErr != nil{
		fmt.Printf("Shutdown error %v", srvErr)
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
	if newNumbers.Second == 0{
		w.WriteHeader(400)
		w.Write([]byte("Error, second number couldn't be zero"))
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
	for i,_:=range allNumbers{
		if i < 2{
			continue
		}
		allNumbers[i] = allNumbers[i-1] + allNumbers[i-2]
	}
	return  (allNumbers[newNumbers.First:])
}
