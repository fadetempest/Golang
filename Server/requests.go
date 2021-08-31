package main

import (
	"fmt"
	"net/http"
	"time"
)

const numJobs = 3

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
	result:=make(chan string, len(urls))
	urlsChan:= make(chan string, len(urls))

	now:=time.Now()
	client:= http.Client{
		Timeout: 5 * time.Second,
	}

	for w:=0;w<numJobs;w++{
		go workerPool(client,result,urlsChan)
	}

	for _, url:=range urls{
		urlsChan<-url
	}
	close(urlsChan)

	for i:=0;i<len(urls);i++{
		fmt.Println(<-result)
	}
	close(result)
	fmt.Println(time.Since(now))
}

func workerPool(client http.Client, result,urlsChan chan string){
	for j:= range urlsChan{
		resp, _:=client.Get(j)
		if resp != nil{
			result<-resp.Status
		} else {
			result<-"404 Not found"
		}
	}
}
