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
	getResp(urls)
}

func getResp(urls []string) {
	client:= http.Client{
		Timeout: 5 * time.Second,
	}
	for _, url:=range urls{
		resp, _:=client.Get(url)
		if resp != nil{
			fmt.Println(resp.Status)
		} else {
			fmt.Println("404 Bad request")
		}
	}
}