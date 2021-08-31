package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}
	client:= http.Client{
		Timeout: 5 * time.Second,
	}
	result:=make(chan string, len(urls))
	now:=time.Now()
	for _, url:=range urls{
		go workerPool(url,client,result)
	}
	for i:=0;i<len(urls);i++{
		fmt.Println(<-result)
	}
	close(result)
	fmt.Println(time.Since(now))
}

func workerPool(url string, client http.Client,result chan string){
	resp, _:=client.Get(url)
	if resp != nil{
		result<-resp.Status
	} else {
		result<-"404 Not found"
	}
}
